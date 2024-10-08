package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 尽管ctx会过期，但是在任何情况下调用它的cancel函数都是很好的实践
	// 如果不这样做，可能会使上下文及父类存活的时间超过必要的时间。
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(time.Second):
		fmt.Println("one second has gone")
	}
}
