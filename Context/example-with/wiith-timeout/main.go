package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			fmt.Println("connect timeout")
			break LOOP
		default:
		}
	}
	fmt.Println("worker Done!")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50) // 设置50ms过期
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
