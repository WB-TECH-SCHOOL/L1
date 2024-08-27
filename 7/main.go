package main

import (
	"errors"
	"fmt"
	"sync"
)

type Cache interface {
	Set(key int, value any)
	Get(key int) (any, error)
}

type cache struct {
	cache map[int]any
	mu    sync.RWMutex
}

func InitCache() Cache {
	return &cache{
		cache: make(map[int]any),
		mu:    sync.RWMutex{},
	}
}

func (c *cache) Set(key int, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

func (c *cache) Get(key int) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, ok := c.cache[key]
	if ok {
		return data, nil
	}
	return nil, errors.New("key not found")
}

// Пример с конкурентной записью и чтением, некоторые ключи еще не имеются в мапе, но мы их запрашиваем
// P.S. тут не только конкурентная запись, но и чтение
func main() {
	c := InitCache()

	var wg sync.WaitGroup
	wg.Add(20)

	for i := range 10 {
		go func(c Cache, i int) {
			defer wg.Done()
			for j := i * 10; j < (i+1)*10; j++ {
				c.Set(j, j)
			}
		}(c, i)
	}

	for i := range 10 {
		go func(c Cache, i int) {
			defer wg.Done()
			for j := i * 10; j < (i+1)*10; j++ {
				data, err := c.Get(j)
				if err != nil {
					fmt.Println(err, j)
				} else {
					fmt.Println(data)
				}
			}
		}(c, i)
	}

	wg.Wait()
}
