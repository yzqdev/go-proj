package grate

import (
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	minWindowTime = int64(time.Millisecond * 500)
	maxWindowTime = int64(time.Millisecond * 2000)

	minLimit = 8
	maxLimit = 2048
)

// VegasStats is the Statistics of vegas.
type VegasStats struct {
	Limit    int64
	InFlight int64
	MinRTT   time.Duration
	LastRTT  time.Duration
}

// VegasLimiter tcp vegas.
type Vegas struct {
	limit      int64
	inFlight   int64
	updateTime int64
	minRTT     int64

	sample atomic.Value
	mu     sync.Mutex
	probes int64
}

// NewVegas new a rate vegas.
func NewVegas() *Vegas {
	v := &Vegas{
		probes: 100,
		limit:  minLimit,
	}
	v.sample.Store(&sample{})
	return v
}

// Stats return the statistics of vegas.
func (v *Vegas) Stats() VegasStats {
	return VegasStats{
		Limit:    atomic.LoadInt64(&v.limit),
		InFlight: atomic.LoadInt64(&v.inFlight),
		MinRTT:   time.Duration(atomic.LoadInt64(&v.minRTT)),
		LastRTT:  time.Duration(v.sample.Load().(*sample).RTT()),
	}
}

// Acquire No matter success or not,done() must be called at last.
func (v *Vegas) Acquire() (done func(time.Time, Operation), success bool) {
	inFlight := atomic.AddInt64(&v.inFlight, 1)
	if inFlight <= atomic.LoadInt64(&v.limit) {
		success = true
	}

	return func(start time.Time, op Operation) {
		atomic.AddInt64(&v.inFlight, -1)
		if op == Ignore {
			return
		}

		end := time.Now().UnixNano()
		rtt := end - start.UnixNano()

		s := v.sample.Load().(*sample)
		v.addToSample(s, rtt, inFlight, op)

		if end > atomic.LoadInt64(&v.updateTime) && s.Count() >= 16 {
			v.mu.Lock()
			defer v.mu.Unlock()

			if v.sample.Load().(*sample) != s {
				return
			}
			v.sample.Store(&sample{})

			lastRTT := s.RTT()
			if lastRTT <= 0 {
				return
			}

			limit := atomic.LoadInt64(&v.limit)
			v.newUpdateTime(s, lastRTT, end)
			v.newMinRTT(s, lastRTT, limit)
			v.newLimit(s, lastRTT, limit)
		}
	}, success
}

func (v *Vegas) addToSample(s *sample, rtt int64, inFlight int64, op Operation) {
	if op == Drop {
		s.Add(rtt, inFlight, true)
	} else if op == Success {
		s.Add(rtt, inFlight, false)
	}
}

func (v *Vegas) newUpdateTime(s *sample, lastRTT int64, end int64) {
	updateTime := end + lastRTT*5
	if lastRTT*5 < minWindowTime {
		updateTime = end + minWindowTime
	} else if lastRTT*5 > maxWindowTime {
		updateTime = end + maxWindowTime
	}
	atomic.StoreInt64(&v.updateTime, updateTime)
}

func (v *Vegas) newMinRTT(s *sample, lastRTT int64, limit int64) {
	v.probes--
	if v.probes <= 0 {
		maxFlight := s.MaxInFlight()
		if maxFlight*2 < v.limit || maxFlight <= minLimit {
			v.probes = 3*limit + rand.Int63n(3*limit)
			v.minRTT = lastRTT
		}
	}

	if v.minRTT == 0 || lastRTT < v.minRTT {
		v.minRTT = lastRTT
	}
}

func (v *Vegas) newLimit(s *sample, lastRTT int64, limit int64) {
	queue := float64(limit) * (1 - float64(v.minRTT)/float64(lastRTT))

	var newLimit float64
	threshold := math.Sqrt(float64(limit)) / 2
	if s.Drop() {
		newLimit = float64(limit) - threshold
		v.storeLimit(newLimit)
		return
	}

	if s.MaxInFlight()*2 < v.limit {
		return
	}

	if queue < threshold {
		newLimit = float64(limit) + 6*threshold
	} else if queue < 2*threshold {
		newLimit = float64(limit) + 3*threshold
	} else if queue < 3*threshold {
		newLimit = float64(limit) + threshold
	} else if queue > 6*threshold {
		newLimit = float64(limit) - threshold
	} else {
		return
	}

	v.storeLimit(newLimit)
}

func (v *Vegas) storeLimit(newLimit float64) {
	newLimit = math.Max(minLimit, math.Min(maxLimit, newLimit))
	atomic.StoreInt64(&v.limit, int64(newLimit))
}
