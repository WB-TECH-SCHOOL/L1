package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func withChan(escape <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-escape:
			fmt.Println("Горутина с каналом завершилась")
			return
		default:
			fmt.Println("Горутина с каналом работает")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func withCtxCancel(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина с контекстом завершилась")
			return
		default:
			fmt.Println("Горутина с контекстом работает")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func withCtxDeadline(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина с дедлайном завершилась")
			return
		default:
			fmt.Println("Горутина с дедлайном работает")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func withCtxTimeout(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина с таймаутом завершилась")
			return
		default:
			fmt.Println("Горутина с таймаутом работает")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	wg.Add(1)
	escape := make(chan struct{})
	go withChan(escape)
	time.Sleep(time.Second * 2)
	escape <- struct{}{}
	wg.Wait()

	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go withCtxCancel(ctx)
	time.Sleep(time.Second * 2)
	cancel()
	wg.Wait()

	wg.Add(1)
	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()
	go withCtxDeadline(ctx)
	time.Sleep(time.Second * 3)
	wg.Wait()

	wg.Add(1)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go withCtxTimeout(ctx)
	time.Sleep(time.Second * 3)
	wg.Wait()
}
