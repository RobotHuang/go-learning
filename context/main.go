package main

import (
	"context"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	defer wg.Done()
Loop:
	for {
		println("work")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break Loop
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	println("over")
}
