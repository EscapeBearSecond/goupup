package main

import (
	"fmt"
	"sync"
	"time"
)

// 扇入扇出模式
// 扇出：多个函数可以从同一个管道中读取数据
// 扇入：一个函数可以从多个管道中读取数据
func main() {
	start := time.Now().UnixMilli()
	buy := buy(10000000000)
	build1 := build(buy)
	build2 := build(buy)
	build3 := build(buy)

	merge := merge(build1, build3, build2)
	packs := pack(merge)
	for p := range packs {
		fmt.Println(p)
	}
	end := time.Now().UnixMilli()
	fmt.Println("时间:", end-start)
}
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	for _, cs := range ins {
		go p(cs)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

func build(in <-chan string) <-chan string {
	time.Sleep(time.Millisecond * 10)
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}
