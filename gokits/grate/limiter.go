package grate

import (
	"context"
	"time"
)

// Operation operations type.
type Operation int

type LogStatsFunc func()

const (
	// Success operation type: success
	Success Operation = iota
	// Ignore operation type: ignore
	Ignore
	// Drop operation type: drop
	Drop
)

// Limiter limit interface.
type Limiter interface {
	Allow(ctx context.Context) (func(Operation), error)
}

// limiter use tcp vegas + codel for adaptive limit.
type limiter struct {
	rate  *Vegas
	queue *CoDel
}

// NewLimiter returns a new Limiter that allows events up to adaptive rtt.
func NewLimiter(c ...CoDelOption) Limiter {
	l := &limiter{
		rate:  NewVegas(),
		queue: NewCoDel(c...),
	}

	return l
}

// Allow implements Limiter.
// if error is returned, no need to call done()
func (l *limiter) Allow(ctx context.Context) (func(op Operation), error) {
	var (
		done func(time.Time, Operation)
		err  error
		ok   bool
	)

	if done, ok = l.rate.Acquire(); !ok {
		// NOTE exceed max inflight, use queue
		if err = l.queue.Push(ctx); err != nil {
			done(time.Time{}, Ignore)
			return func(op Operation) {}, err
		}
	}

	start := time.Now()
	return func(op Operation) {
		done(start, op)
		l.queue.Pop()
	}, nil
}

// Stats return vegas and queue statistics
func (l *limiter) Stats() (VegasStats, CoDelStats) {
	return l.rate.Stats(), l.queue.Stats()
}
