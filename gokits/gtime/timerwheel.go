package gtime

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"sync"
	"time"

	"gitee.com/yzqdev/gokits/grand"
)

const (
	defaultTickDuration = 100 * time.Millisecond
	defaultWheelCount   = 512
)

// TimerWheel 时钟轮
type TimerWheel struct {
	lock          sync.Mutex    // 锁
	tickDuration  time.Duration // 卡槽每次跳动的时间间隔
	roundDuration time.Duration // 一轮耗时
	wheelCount    int           // 卡槽数
	wheel         []*iterator   // 卡槽
	quit          chan struct{} // 退出
	wheelCursor   int           // 当前卡槽位置
}

// iterator is a iterator for wheel timeout
type iterator struct {
	items map[string]*wheelTimeOut
}

// wheelTimeOut is a object to process timeout event
type wheelTimeOut struct {
	id       string        // 定时器标识
	delay    time.Duration // 延迟时间
	index    int           // 卡槽索引位置
	rounds   int           // 需要转动的周期数
	task     func()        // 到期执行的任务
	outCh    chan struct{} // 到期事件队列
	times    int           // 超时设定的次数
	expTimes int           // 已超时的次数
}

// TwOption is the configuration of TimerWheel
type TwOption struct {
	TickDuration time.Duration `json:"tick_duration" yaml:"tick_duration"`
	WheelCount   int           `json:"wheel_count" yaml:"wheel_count"`
}

// NewTimerWheel create a instance of TimerWheel with given options
func NewTimerWheel(opt ...TwOption) *TimerWheel {
	tw := &TimerWheel{
		tickDuration:  defaultTickDuration,
		wheelCount:    defaultWheelCount,
		quit:          make(chan struct{}),
		wheelCursor:   0,
		roundDuration: defaultTickDuration * defaultWheelCount,
	}

	if len(opt) >= 1 {
		tw.tickDuration = opt[0].TickDuration
		tw.wheelCount = opt[0].WheelCount
	}

	tw.roundDuration = tw.tickDuration * time.Duration(tw.wheelCount)
	tw.createWheel(tw.wheelCount)
	return tw
}

// Start a ticker for check expired items
func (t *TimerWheel) Start() {
	tick := time.NewTicker(t.tickDuration)
	go func() {
		for {
			select {
			case <-tick.C:
				t.wheelCursor++
				if t.wheelCursor == t.wheelCount {
					t.wheelCursor = 0
				}

				iterator := t.wheel[t.wheelCursor]
				tasks := t.fetchExpiredTimeouts(iterator)
				t.notifyExpiredTimeOut(tasks)
			case <-t.quit:
				tick.Stop()
				return
			}
		}
	}()
}

// Stop the ticker
func (t *TimerWheel) Stop() {
	t.quit <- struct{}{}
}

func (t *TimerWheel) createWheel(wheelCount int) {
	arr := make([]*iterator, wheelCount)
	for v := 0; v < wheelCount; v++ {
		arr[v] = &iterator{items: make(map[string]*wheelTimeOut)}
	}
	t.wheel = arr
}

// AfterFunc add a timer callback function which will trigger after the given interval time and trigger times
func (t *TimerWheel) AfterFunc(interval time.Duration, times int, f func()) (string, error) {
	if f == nil {
		return "", errors.New("timer callback function is empty")
	}

	if interval <= 0 {
		return "", errors.New("interval Must be greater than zero")
	}

	if times <= 0 {
		times = 0x0FFFFFFF
	}

	timeOut := &wheelTimeOut{
		delay: interval,
		task:  f,
		times: times,
	}

	tid := t.scheduleTimeOut(timeOut)
	return tid, nil
}

// After add a timer using the given interval time
func (t *TimerWheel) After(interval time.Duration, times int) (chan struct{}, string, error) {
	if interval <= 0 {
		return nil, "", errors.New("interval Must be greater than zero")
	}

	if times <= 0 {
		times = 0x0FFFFFFF
	}

	timeOut := &wheelTimeOut{
		delay: interval,
		outCh: make(chan struct{}),
		times: times,
	}

	tid := t.scheduleTimeOut(timeOut)
	return timeOut.outCh, tid, nil
}

// Cancel the timer
func (t *TimerWheel) Cancel(timerID string) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, it := range t.wheel {
		delete(it.items, timerID)
	}
	return nil
}

func (t *TimerWheel) scheduleTimeOut(timeOut *wheelTimeOut) string {
	if timeOut.delay < t.tickDuration {
		timeOut.delay = t.tickDuration
	}
	lastRoundDelay := timeOut.delay % t.roundDuration
	lastTickDelay := timeOut.delay % t.tickDuration

	// 计算卡槽位置
	relativeIndex := lastRoundDelay / t.tickDuration
	if lastTickDelay != 0 {
		relativeIndex = relativeIndex + 1
	}
	// 计算时间轮圈数
	remainingRounds := timeOut.delay / t.roundDuration
	if lastRoundDelay == 0 {
		remainingRounds = remainingRounds - 1
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	stopIndex := t.wheelCursor + int(relativeIndex)
	if stopIndex >= t.wheelCount {
		stopIndex = stopIndex - t.wheelCount
		timeOut.rounds = int(remainingRounds) + 1
	} else {
		timeOut.rounds = int(remainingRounds)
	}
	timeOut.index = stopIndex
	item := t.wheel[stopIndex]
	if item == nil {
		item = &iterator{
			items: make(map[string]*wheelTimeOut),
		}
	}

	if timeOut.id == "" {
		timeOut.id = t.guid()
	}

	item.items[timeOut.id] = timeOut
	t.wheel[stopIndex] = item

	return timeOut.id
}

// 判断当前卡槽中是否有超时任务,将超时task加入切片中
func (t *TimerWheel) fetchExpiredTimeouts(iter *iterator) []*wheelTimeOut {
	t.lock.Lock()
	defer t.lock.Unlock()

	task := []*wheelTimeOut{}

	for k, v := range iter.items {
		if v.rounds <= 0 { //已经超时了
			task = append(task, v)
			delete(iter.items, k)
		} else {
			v.rounds--
		}
	}

	return task
}

// 执行超时任务
func (t *TimerWheel) notifyExpiredTimeOut(tasks []*wheelTimeOut) {
	for _, task := range tasks {
		task.expTimes++
		if task.expTimes < task.times { // 如果执行的次数小于设置的次数，则再次调度
			_ = t.scheduleTimeOut(task)
		}

		if task.task != nil {
			go task.task()
		} else {
			go func(times int) {
				task.outCh <- struct{}{}
				if times == task.times {
					close(task.outCh)
				}
			}(task.expTimes)
		}
	}
}

func (t *TimerWheel) md5str(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))

}

func (t *TimerWheel) guid() string {
	return grand.NewUUID()
}
