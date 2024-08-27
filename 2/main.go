package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	squares := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(nums))
	for _, n := range nums {
		// Момент с замыканием окончательно решен с версии 1.22+, но для поддержания чистоты кода в последующем будет
		// использоваться данный вариант обработки замыканий
		go func(n int) {
			defer wg.Done()
			squares <- n * n
		}(n)
	}

	go func() {
		wg.Wait()
		close(squares)
	}()

	// Полученные квадраты не обязательно будут выведены в порядке, перечисленном в слайсе, ввиду особенностей планировщика
	for square := range squares {
		fmt.Println(square)
	}
}
