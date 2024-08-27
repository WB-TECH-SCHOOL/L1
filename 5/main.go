package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// --------------------- ЧЕРЕЗ ТАЙМЕР ---------------------
//func main() {
//	var N int
//	fmt.Scan(&N)
//
//	timer := time.NewTimer(time.Second * time.Duration(N))
//	dataCh := make(chan any)
//
//	go func() {
//		for v := range dataCh {
//			fmt.Println(v)
//		}
//	}()
//
//	for {
//		select {
//		case <-timer.C:
//			return
//		default:
//			dataCh <- rand.Int()
//			time.Sleep(time.Millisecond * 100)
//		}
//	}
//}

// --------------------- ЧЕРЕЗ КОНТЕКСТ ---------------------
func main() {
	var N int
	fmt.Scan(&N)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(N))
	defer cancel()
	dataCh := make(chan any)

	var wg sync.WaitGroup
	wg.Add(2)
	// Пишущая горутина
	go func(ctx context.Context, dataCh chan<- any) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				close(dataCh)
				return
			default:
				dataCh <- rand.Int()
				time.Sleep(time.Millisecond * 200)
			}
		}
	}(ctx, dataCh)

	// Читающая горутина
	go func(ctx context.Context, dataCh <-chan any) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case data, ok := <-dataCh:
				if ok {
					fmt.Println(data)
				}
			}
		}
	}(ctx, dataCh)

	wg.Wait()
}
