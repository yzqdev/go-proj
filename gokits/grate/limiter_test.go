package grate

import (
	"context"
	"testing"
	"time"
)

func TestRateSuccess(t *testing.T) {
	ch := make(chan struct{})
	go worker(100, ch)
	failed := producerL(NewLimiter(), 100, ch)
	if failed > 0 {
		t.Fatalf("Should be rejected 0 time,but (%d)", failed)
	}
}

func TestRateFail(t *testing.T) {
	ch := make(chan struct{})
	go worker(100, ch)
	failed := producerL(NewLimiter(), 200, ch)
	if failed < 900 {
		t.Fatalf("Should be rejected more than 900 times,but (%d)", failed)
	}
}

func TestRateFailMuch(t *testing.T) {
	ch := make(chan struct{})
	go worker(10, ch)
	failed := producerL(NewLimiter(), 200, ch)
	if failed < 1600 {
		t.Fatalf("Should be rejected more than 1600 times,but (%d)", failed)
	}
}

func producerL(l Limiter, qps int64, ch chan struct{}) (failed int) {
	for i := 0; i < int(qps)*10; i++ {
		go func() {
			done, err := l.Allow(context.Background())
			if err != nil {
				failed++
				return
			}

			defer done(Success)
			ch <- struct{}{}
		}()
		time.Sleep(time.Duration(int64(time.Second) / qps))
	}
	return
}
