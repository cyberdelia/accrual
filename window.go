package accrual

import (
	"container/ring"
	"sync"
	"time"
)

type window struct {
	r *ring.Ring
	l time.Time
	m sync.RWMutex
}

// NewMemoryWindow creates a new in-memory window.
func NewMemoryWindow(size int) Window {
	return &window{
		r: ring.New(size),
	}
}

// Record records a new heartbeat.
func (w *window) Record() {
	w.m.Lock()
	defer w.m.Unlock()
	now := time.Now()
	if !w.l.IsZero() {
		w.r.Value = now.Sub(w.l).Nanoseconds()
		w.r = w.r.Next()
	}
	w.l = now
}

// Last returns the last heartbeat time.
func (w *window) Last() time.Time {
	return w.l
}

// Distribution returns all recorded intervals.
func (w *window) Distribution() []int64 {
	w.m.RLock()
	defer w.m.RUnlock()
	var values []int64
	w.r.Do(func(v interface{}) {
		if i, ok := v.(int64); ok {
			values = append(values, i)
		}
	})
	return values
}
