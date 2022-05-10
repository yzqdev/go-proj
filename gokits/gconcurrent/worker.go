package gconcurrent

import (
	"context"
	"errors"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrTimeout = errors.New("add to job queue timeout")
)

// JobFunc is a job function will execute by worker goroutine
type JobFunc func(context.Context)

// PanicFunc is a function will process panic which throw in worker function
type PanicFunc func(recovered interface{}, funcName string)

type queueItem struct {
	jobFunc  JobFunc
	funcName string
}

type WorkerPool interface {
	// Execute add worker function to queue to wait for executing by worker goroutine
	Execute(f JobFunc)

	// Execute summit worker function to queue  with a timeout timer,
	// if add queue success it will wait for executing by worker goroutine
	Submit(f JobFunc, timeout time.Duration) error

	// Stop cancel all goroutines started by this pool and wait
	Shutdown(ctx context.Context)

	// Stats return the statistics of worker pool
	Stats() WpStats

	// Option return the configuration of worker pool
	Option() WpOption
}

type workerPool struct {
	queue  chan *queueItem
	option WpOption
	stats  WpStats

	mux    sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
}

// WpOption is the worker pool parameter
type WpOption struct {
	InitWorkerNum int       `json:"init_worker_num" yaml:"init_worker_num"`
	MaxWorkerNum  int       `json:"max_worker_num" yaml:"max_worker_num"`
	QueueSize     int       `json:"queue_size" yaml:"queue_size"`
	PanicFunc     PanicFunc `json:"-"`
}

// WpStats is the statistics of worker pool
type WpStats struct {
	ActiveNum     int32
	WorkerNum     int32
	ExecuteNum    int32
	SubmitFailNum int32
	PanicNum      int32
}

// NewWorkerPool creates a instance of WorkerPool with given option
// default parameters:
//   queueSize: 100
//   initWorkerNum: 2
//   maxWorkerNum: 50
func NewWorkerPool(opt ...WpOption) WorkerPool {
	ctx, cancel := context.WithCancel(context.TODO())
	wp := &workerPool{
		option: WpOption{
			InitWorkerNum: 2,
			MaxWorkerNum:  50,
			QueueSize:     100,
		},
		ctx:    ctx,
		cancel: cancel,
	}

	if len(opt) >= 1 {
		cfg := opt[0]
		wp.option = cfg
	}

	wp.queue = make(chan *queueItem, wp.option.QueueSize)
	wp.run(wp.option.InitWorkerNum)

	return wp
}

func (w *workerPool) Stats() WpStats {
	return w.stats
}

func (w *workerPool) Option() WpOption {
	return w.option
}

func (w *workerPool) toItem(jf JobFunc) *queueItem {
	pc := reflect.ValueOf(jf).Pointer()
	f := runtime.FuncForPC(pc)
	funcName := f.Name()

	return &queueItem{
		jobFunc:  jf,
		funcName: funcName,
	}
}

func (w *workerPool) run(incNum int) {
	for idx := 0; idx < incNum; idx++ {
		atomic.AddInt32(&w.stats.WorkerNum, 1)
		go func() {
			for it := range w.queue {
				w.executeOne(it)
			}
		}()
	}
}

func (w *workerPool) executeOne(it *queueItem) {
	atomic.AddInt32(&w.stats.ActiveNum, 1)

	defer func() {
		atomic.AddInt32(&w.stats.ActiveNum, -1)
		recovered := recover()
		if recovered != nil {
			atomic.AddInt32(&w.stats.PanicNum, 1)
			if w.option.PanicFunc != nil {
				w.option.PanicFunc(recovered, it.funcName)
			}
		}
	}()

	atomic.AddInt32(&w.stats.ExecuteNum, 1)
	it.jobFunc(w.ctx)
}

func (w *workerPool) incWorker() {
	activeNum := int(atomic.LoadInt32(&w.stats.ActiveNum))
	workerNum := int(atomic.LoadInt32(&w.stats.WorkerNum))
	if activeNum == workerNum && workerNum < w.option.MaxWorkerNum {
		w.mux.Lock()
		atomic.LoadInt32(&w.stats.WorkerNum)
		incNum := workerNum / 2
		if incNum < 1 {
			incNum = 1
		}
		if workerNum+incNum > w.option.MaxWorkerNum {
			incNum = w.option.MaxWorkerNum - int(workerNum)
		}
		w.run(int(incNum))
		w.mux.Unlock()
	}
}

func (w *workerPool) Execute(jf JobFunc) {
	w.incWorker()
	w.queue <- w.toItem(jf)
}

func (w *workerPool) Submit(jf JobFunc, timeout time.Duration) error {
	w.incWorker()
	select {
	case <-time.NewTimer(timeout).C:
		atomic.AddInt32(&w.stats.SubmitFailNum, 1)
		return ErrTimeout
	case w.queue <- w.toItem(jf):
		return nil
	}
}

func (w *workerPool) Shutdown(ctx context.Context) {
	close(w.queue)
	w.queue = nil
	w.cancel()

	for {
		timer := time.NewTimer(time.Millisecond * 100)
		select {
		case <-timer.C:
			if w.checkNoActive() {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func (w *workerPool) checkNoActive() bool {
	return atomic.LoadInt32(&w.stats.ActiveNum) == 0
}
