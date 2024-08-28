package main

import (
	"context"
	"time"
)

func sleep(d time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	<-ctx.Done()
}

func main() {
	sleep(time.Second * 2)
}
