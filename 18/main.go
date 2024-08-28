package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type counter struct {
	count atomic.Uint64
}

func initCounter() *counter {
	return &counter{count: atomic.Uint64{}}
}

func (c *counter) inc() {
	c.count.Add(1)
}

func (c *counter) get() uint64 {
	return c.count.Load()
}

func main() {
	c := counter{}

	for _ = range 1000 {
		go c.inc()
	}

	time.Sleep(time.Second)
	fmt.Println(c.get())
}
