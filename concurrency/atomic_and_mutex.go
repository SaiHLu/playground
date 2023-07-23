package main

import (
	"sync/atomic"
)

type State struct {
	/** Mutext Way
	mu    sync.Mutex
	 count int
	**/

	/** Atomic way (only use for simple process, e.g. counter, boolean) **/
	count int32
}

func (s *State) setState(i int) {
	// s.mu.Lock()
	// defer s.mu.Unlock()
	// s.count = i
	atomic.AddInt32(&s.count, int32(i))
}

func atomicAndMutex() {
}
