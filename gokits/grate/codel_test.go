package grate

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	testConf = CoDelOption{
		Target:   20,
		Internal: 500,
	}
	qps = time.Microsecond * 2000
)

func testPush(q *CoDel, sleep time.Duration, delay time.Duration, drop *int64, tm *int64) {
	var group sync.WaitGroup
	for i := 0; i < 5000; i++ {
		time.Sleep(sleep)
		group.Add(1)
		go func() {
			defer group.Done()
			ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*1000))
			defer cancel()
			if err := q.Push(ctx); err != nil {
				if err == ErrLimitExceed {
					atomic.AddInt64(drop, 1)
				} else {
					atomic.AddInt64(tm, 1)
				}
			} else {
				time.Sleep(delay)
				q.Pop()
			}
		}()
	}
	group.Wait()
}

func TestCoDel3000(t *testing.T) {
	q := NewCoDel(testConf)
	drop := new(int64)
	tm := new(int64)
	delay := time.Millisecond * 3000
	testPush(q, qps, delay, drop, tm)
	fmt.Printf("qps %v process time %v drop %d timeout %d \n", int64(time.Second/qps), delay, *drop, *tm)
	time.Sleep(time.Second)
}

func TestCoDel2000(t *testing.T) {
	q := NewCoDel(testConf)
	drop := new(int64)
	tm := new(int64)
	delay := time.Millisecond * 2000
	testPush(q, qps, delay, drop, tm)
	fmt.Printf("qps %v process time %v drop %d timeout %d \n", int64(time.Second/qps), delay, *drop, *tm)
	time.Sleep(time.Second)
}

func TestCoDel1000(t *testing.T) {
	q := NewCoDel(testConf)
	drop := new(int64)
	tm := new(int64)
	delay := time.Millisecond * 1000
	testPush(q, qps, delay, drop, tm)
	fmt.Printf("qps %v process time %v drop %d timeout %d \n", int64(time.Second/qps), delay, *drop, *tm)
}

func TestCoDel500(t *testing.T) {
	q := NewCoDel(testConf)
	drop := new(int64)
	tm := new(int64)
	delay := time.Millisecond * 500
	testPush(q, qps, delay, drop, tm)
	fmt.Printf("qps %v process time %v drop %d timeout %d \n", int64(time.Second/qps), delay, *drop, *tm)
}

func BenchmarkAQM(b *testing.B) {
	q := NewCoDel()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*5))
			err := q.Push(ctx)
			if err == nil {
				q.Pop()
			}
			cancel()
		}
	})
}
