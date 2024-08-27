package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	nums := []uint32{2, 4, 6, 8, 10}

	var ans uint32
	var wg sync.WaitGroup
	wg.Add(len(nums))
	for _, n := range nums {
		go func(n uint32) {
			defer wg.Done()
			// Можно через мютекс, но нет необходимости, ибо операция слишком проста
			atomic.AddUint32(&ans, n*n)
		}(n)
	}

	wg.Wait()

	fmt.Println(ans)
}
