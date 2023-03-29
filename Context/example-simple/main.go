package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(3)
	go work(ctx, "node1")
	go work(ctx, "node2")
	go work(ctx, "node3")
	time.Sleep(5 * time.Second)
	fmt.Println("stop the goroutine")
	cancel()
	wg.Wait()
}
func work(ctx context.Context, node string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(node, "got the stop channel")
			wg.Done()
			return
		default:
			fmt.Println(node, "still working")
			if node == "node1" {
				time.Sleep(1 * time.Second)
			} else if node == "node2" {
				time.Sleep(2 * time.Second)
			} else if node == "node3" {
				time.Sleep(3 * time.Second)
			}
		}
	}
}
