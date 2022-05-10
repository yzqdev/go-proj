package internal

import (
	"testing"
	"time"
	"unsafe"
)

func TestGetPut(t *testing.T) {
	wrm := NewTimeMap()
	k := "a"
	kp := uintptr(unsafe.Pointer(&k))
	if wr, ok := wrm.Get(kp); (wr != nil) || (ok) {
		t.Error("Get of non-existing key failed.")
	}
	v := wrm.Put(kp, "result", 5*time.Second)
	if v, ok := v.(string); (!ok) || (v != "result") {
		t.Error("Put error.")
	}
	if wr, ok := wrm.Get(kp); (wr.(string) != "result") || (!ok) {
		t.Error("Get of existing key failed.")
	}
	wrm.Delete(kp)
	if wr, ok := wrm.Get(kp); (wr != nil) || (ok) {
		t.Error("Get of non-existing key failed.")
	}
}

func TestTime(t *testing.T) {
	wrm := NewTimeMap()
	k := "a"
	kp := uintptr(unsafe.Pointer(&k))

	if _, ok := wrm.Get(kp); ok {
		t.Error("Get of non-existing key failed.")
	}
	v := wrm.Put(kp, "result", 5*time.Millisecond)
	if v, ok := v.(string); (!ok) || (v != "result") {
		t.Error("Put error.")
	}
	time.Sleep(time.Second * 6)
	if _, ok := wrm.Get(kp); ok {
		t.Error("Get of existing key failed.")
	}
}
