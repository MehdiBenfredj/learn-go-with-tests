package sync

import "sync"

type Counter struct {
	val int
	mu  *sync.RWMutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}
