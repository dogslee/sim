package sim

import (
	"runtime"
	"sync/atomic"
)

type Mutex struct {
	state int32
}

func (m *Mutex) Lock() {
	for !atomic.CompareAndSwapInt32(&m.state, 0, 1) {
		runtime.Gosched()
	}
}

func (m *Mutex) Unlock() {
	atomic.AddInt32(&m.state, -1)
}
