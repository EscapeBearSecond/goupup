package _case

import (
	"context"
	"fmt"
	"time"
)

func ContextCase() {
	done := make(chan struct{})
	go f1(done)
	go f1(done)
	time.Sleep(time.Second * 2)
	close(done)
}

func f1(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("goroutine exit")
			return
		}
	}
}

func ContextCase2() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "desc", "ContextCase2")
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	ch := make(chan []int)
	data := [][]int{{1, 2}, {2, 3}}
	go calculate(ctx, ch)
	for i := 0; i < len(data); i++ {
		ch <- data[i]
	}
	time.Sleep(time.Second * 10)
}
func calculate(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			// fmt.Println(item)
			ctx := context.WithValue(ctx, "desc", "calculate")

			// 这里重新创建channel的目的是为了防止漏计算
			// 在sumContext中会取出一个值，这样在multiContext中就取不到了
			ch := make(chan []int)
			go sumContext(ctx, ch)
			ch <- item

			ch1 := make(chan []int)
			go multiContext(ctx, ch1)
			ch1 <- item
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("calculate goroutine exit, context desc: %s, error message: %s\n", desc, ctx.Err())
			return
		}
	}
}
func sumContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := sum(a, b)
			fmt.Printf("%d + %d = %d\n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("sumContext goroutine exit, context desc: %s, error message: %s\n", desc, ctx.Err())
			return
		}
	}
}
func multiContext(ctx context.Context, data <-chan []int) {
	for {
		select {
		case item := <-data:
			a, b := item[0], item[1]
			res := multi(a, b)
			fmt.Printf("%d * %d = %d\n", a, b, res)
		case <-ctx.Done():
			desc := ctx.Value("desc").(string)
			fmt.Printf("multiContext goroutine exit, context desc: %s, error message: %s\n", desc, ctx.Err())
			return
		}
	}
}

func sum(a, b int) int {
	return a + b
}
func multi(a, b int) int {
	return a * b
}
