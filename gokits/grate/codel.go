package grate

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"
)

var (
	ErrLimitExceed = errors.New("limit exceed")
	ErrDeadline    = errors.New("deal line")
)

// CoDelOption CoDel queue config.
type CoDelOption struct {
	Target   int64 // target queue delay (default 20 ms).
	Internal int64 // sliding minimum time window width (default 500 ms)
}

// CoDelStats is the Statistics of CoDel queue.
type CoDelStats struct {
	Dropping bool
	FaTime   int64
	DropNext int64
	Packets  int
}

type queuePacket struct {
	dropChan  chan bool
	timestamp int64
}

// CoDel is CoDel req buffer queue.
type CoDel struct {
	pool     sync.Pool
	packets  chan queuePacket
	mux      sync.RWMutex
	option   *CoDelOption
	count    int64
	dropping bool  // in drop state
	faTime   int64 // Time when we'll declare we're above target (0 if below)
	dropNext int64 // Packets dropped since going into drop state
}

// NewCoDel creates a CoDel queue.
func NewCoDel(opt ...CoDelOption) *CoDel {
	option := CoDelOption{
		Target:   50,
		Internal: 500,
	}
	if len(opt) >= 1 {
		option = opt[0]
	}

	q := &CoDel{
		packets: make(chan queuePacket, 2048),
		option:  &option,
	}

	q.pool.New = func() interface{} {
		return make(chan bool)
	}
	return q
}

// Stats return the statistics of codel
func (q *CoDel) Stats() CoDelStats {
	q.mux.Lock()
	defer q.mux.Unlock()

	return CoDelStats{
		Dropping: q.dropping,
		FaTime:   q.faTime,
		DropNext: q.dropNext,
		Packets:  len(q.packets),
	}
}

// Push req into CoDel request buffer queue.
// if return error is nil,the caller must call q.Done() after finish request handling
func (q *CoDel) Push(ctx context.Context) (err error) {
	r := queuePacket{
		dropChan:  q.pool.Get().(chan bool),
		timestamp: time.Now().UnixNano() / int64(time.Millisecond),
	}

	select {
	case q.packets <- r:
	default:
		err = ErrLimitExceed
		q.pool.Put(r.dropChan)
	}

	if err != nil {
		return
	}

	select {
	case drop := <-r.dropChan:
		if drop {
			err = ErrLimitExceed
		}
		q.pool.Put(r.dropChan)
	case <-ctx.Done():
		err = ErrDeadline
	}
	return
}

// Pop req from CoDel request buffer queue.
func (q *CoDel) Pop() {
	for {
		select {
		case p := <-q.packets:
			drop := q.judge(p)
			select {
			case p.dropChan <- drop:
				if !drop {
					return
				}
			default:
				q.pool.Put(p.dropChan)
			}
		default:
			return
		}
	}
}

func (q *CoDel) controlLaw(now int64) int64 {
	q.dropNext = now + int64(float64(q.option.Internal)/math.Sqrt(float64(q.count)))
	return q.dropNext
}

func (q *CoDel) checkFaTime(now, elapsed int64) (drop bool) {
	if elapsed < q.option.Target {
		q.faTime = 0
	} else if q.faTime == 0 {
		q.faTime = now + q.option.Internal
	} else if now >= q.faTime {
		return true
	}
	return false
}

// judge decide if the packet should drop or not.
func (q *CoDel) judge(p queuePacket) (drop bool) {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	elapsed := now - p.timestamp

	q.mux.Lock()
	defer q.mux.Unlock()

	drop = q.checkFaTime(now, elapsed)
	if q.dropping {
		if !drop {
			// elapsed time below target - leave dropping state
			q.dropping = false
			return
		}

		if now > q.dropNext {
			q.count++
			q.dropNext = q.controlLaw(q.dropNext)
			drop = true
			return
		}
	}

	if drop && (now-q.dropNext < q.option.Internal || now-q.faTime >= q.option.Internal) {
		q.dropping = true
		// If we're in a drop cycle, the drop rate that controlled the queue
		// on the last cycle is a good starting point to control it now.
		if now-q.dropNext < q.option.Internal {
			if q.count > 2 {
				q.count = q.count - 2
			} else {
				q.count = 1
			}
		} else {
			q.count = 1
		}
		q.dropNext = q.controlLaw(now)
		drop = true
		return
	}
	return
}
