package counter

import (
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func New(value int) *Counter {
	return &Counter{
		value: value,
	}
}
// get current value of the counter and increment it
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	value := c.value
	c.value++
	return value
}

func (c *Counter) Reset(newValue int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = newValue
}
