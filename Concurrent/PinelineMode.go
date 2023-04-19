package main

import "fmt"

// 管道模式
// 每一道工序的输出就是下一道工序的输入
func main() {
	buy := Buy(6)
	build := Build(buy)
	pack := Pack(build)
	for p := range pack {
		fmt.Println(p)
	}
}
func Buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

func Build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}
func Pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}
