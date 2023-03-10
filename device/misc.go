package device

import "sync/atomic"

/* Atomic Boolean */

const (
	AtomicFalse = int32(iota)
	AtomicTrue
)

type AtomicBool struct {
	int32
}

func (a *AtomicBool) Get() bool {
	return atomic.LoadInt32(&a.int32) == AtomicTrue
}

func (a *AtomicBool) Swap(val bool) bool {
	flag := AtomicFalse
	if val {
		flag = AtomicTrue
	}
	return atomic.SwapInt32(&a.int32, flag) == AtomicTrue
}

func (a *AtomicBool) Set(val bool) {
	flag := AtomicFalse
	if val {
		flag = AtomicTrue
	}
	atomic.StoreInt32(&a.int32, flag)
}

func min(a, b uint) uint {
	if a > b {
		return b
	}
	return a
}
