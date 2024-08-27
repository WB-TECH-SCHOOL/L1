package main

import (
	"fmt"
)

func main() {
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	chanX := make(chan int)
	chan2X := make(chan int)

	// Горутина для отправки данных в канал X
	go func() {
		for _, n := range nums {
			chanX <- n
		}
		close(chanX)
	}()

	// Горутина для обработки данных и отправки в канал Y
	go func() {
		for n := range chanX {
			chan2X <- n * 2
		}
		close(chan2X)
	}()

	// Горутина для чтения из канала Y и вывода результатов
	for n := range chan2X {
		fmt.Println(n)
	}
}
