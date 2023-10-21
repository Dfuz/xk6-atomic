package atomic

import (
	"sync/atomic"
)

type counter struct {
	i int64
}

// Inc increments the counter and returns the new value.
func (c *counter) Inc() int64 {
	return atomic.AddInt64(&c.i, 1)
}

// Dec decrements the counter and returns the new value.
func (c *counter) Dec() int64 {
	return atomic.AddInt64(&c.i, -1)
}

// Add adds i to the counter and returns the new value.
func (c *counter) Add(i int64) int64 {
	return atomic.AddInt64(&c.i, i)
}

// Value returns the current value of the counter.
func (c *counter) Value() int64 {
	return atomic.LoadInt64(&c.i)
}
