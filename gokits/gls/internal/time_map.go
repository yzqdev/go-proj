package internal

import (
	"sync"
	"time"
)

// TimeMap interface
type TimeMap interface {
	// Get
	Get(key uintptr) (val interface{}, ok bool)

	// Put
	Put(key uintptr, obj interface{}, timeout time.Duration) interface{}

	// Delete
	Delete(key uintptr)
}

const bucketConst = 64

type bucket struct {
	objects map[uintptr]*storeValue
	mu      sync.RWMutex
}

type storeValue struct {
	obj      interface{}
	expireAt time.Time
}

type timemap []*bucket

func (tm timemap) getbucket(key uintptr) *bucket {
	return tm[uint(key)%uint(bucketConst)]
}

func (b *bucket) get(key uintptr) (interface{}, bool) {
	v, ok := b.objects[key]
	if !ok {
		return nil, false
	}
	return v.obj, true
}

func (tm timemap) Get(key uintptr) (interface{}, bool) {
	b := tm.getbucket(key)
	b.mu.RLock()
	v, ok := b.get(key)
	b.mu.RUnlock()
	return v, ok
}

func (b *bucket) put(key uintptr, obj interface{}, timeout time.Duration) interface{} {
	vo, ok := b.objects[key]
	if ok {
		return vo.obj
	}
	b.objects[key] = &storeValue{obj: obj, expireAt: time.Now().Add(timeout)}
	return obj
}

func (tm timemap) Put(key uintptr, obj interface{}, timeout time.Duration) interface{} {
	b := tm.getbucket(key)
	b.mu.Lock()
	v := b.put(key, obj, timeout)
	b.mu.Unlock()
	return v
}

func (tm timemap) Delete(key uintptr) {
	b := tm.getbucket(key)
	b.mu.Lock()
	delete(b.objects, key)
	b.mu.Unlock()
}

func (tm timemap) goroutine() {
	t := time.Tick(5 * time.Second)
	for {
		<-t
		now := time.Now()
		for _, b := range tm {
			b.mu.Lock()
			for k, v := range b.objects {
				if v.expireAt.Before(now) {
					delete(b.objects, k)
				}
			}
			b.mu.Unlock()
		}
	}
}

// NewTimeMap create a new TimeMap
func NewTimeMap() TimeMap {
	tm := make(timemap, bucketConst)
	for i := 0; i < bucketConst; i++ {
		tm[i] = &bucket{
			objects: make(map[uintptr]*storeValue),
		}
	}

	go tm.goroutine()
	return tm
}
