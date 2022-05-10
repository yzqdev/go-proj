package gls

import (
	"time"

	"gitee.com/yzqdev/gokits/gls/internal"
)

var (
	gtm = internal.NewTimeMap()
)

// Storage the current goroutine local storage
type Storage map[string]interface{}

func newStorage() Storage {
	return make(map[string]interface{})
}

func getStorage() Storage {
	v, ok := gtm.Get(uintptr(internal.Getg()))
	if !ok {
		return nil
	}

	s, _ := (v).(Storage)
	return s
}

// Get gets the value by key as it exists for the current goroutine.
func Get(key string) interface{} {
	s := getStorage()
	if s != nil {
		return s[key]
	}
	return nil
}

// Set sets the value by key and associates it with the current goroutine.
func Set(key string, value interface{}) {
	s := getStorage()
	if s != nil {
		s[key] = value
	}
}

// Put return a new map instance for current goroutine,
// it will store a global map and delete after `timeout` Duration
func Put(timeout time.Duration) Storage {
	return putStorage(newStorage(), timeout)
}

func putStorage(s Storage, timeout time.Duration) Storage {
	v := gtm.Put(uintptr(internal.Getg()), s, timeout)
	rs, _ := v.(Storage)
	return rs
}

// Cleanup removes all data associated with this goroutine. If this is not
// called, the data may persist for the lifetime of your application. This
// must be called from the very first goroutine to invoke Set
func Cleanup() {
	gtm.Delete(uintptr(internal.Getg()))
}

// With is a convenience function that stores the given values on this
// goroutine, calls the provided function (which will have access to the
// values) and then cleans up after itself.
func With(values Storage, f func()) {
	putStorage(values, 1*time.Hour)
	f()
	Cleanup()
}

func copyStorage(src Storage, dest Storage) {
	for k, v := range src {
		dest[k] = v
	}
}

// Go creates a new goroutine and runs the provided function in that new
// goroutine. It also associates any key,value pairs stored for the parent
// goroutine with the child goroutine. This function must be used if you wish
// to preserve the reference to any data stored in gls. This function
// automatically cleans up after itself. Do not call cleanup in the function
// passed to this function.
func Go(f func()) {
	ps := getStorage()
	cs := newStorage()
	if ps != nil {
		copyStorage(ps, cs)
	}
	go func() {
		putStorage(cs, 1*time.Hour)
		f()
		Cleanup()
	}()
}
