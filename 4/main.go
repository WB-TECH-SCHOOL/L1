package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Введите кол-во воркеров:")
	var n int
	fmt.Scan(&n)

	data := make(chan int)
	for _ = range n {
		go func(data <-chan int) {
			// Можно сделать через бесконечный цикл и select
			// for {
			//    select {
			//    case v := <-data:
			//        fmt.Println(v)
			//    }
			// }
			// Но тут нет необходимости в непрерывном чтении из канала, нужно читать только по мере поступления данных
			for v := range data {
				fmt.Println(v)
			}
		}(data)
	}

	for {
		data <- rand.Int()
		time.Sleep(time.Millisecond * 100)
	}
}
