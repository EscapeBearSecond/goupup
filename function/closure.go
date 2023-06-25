package function

import (
	"fmt"
	"log"
)

func Fibonacci(n int) int {
	if n <= 2 {
		log.Fatalln("please input a number more than 2")
	}
	t := tool()
	var res int
	for i := 0; i < n-2; i++ {
		res = t()
	}
	return res
}

// 斐波那契数列
func tool() func() int {
	var x0, x1, x2 = 0, 1, 0
	return func() int {
		x2 = x0 + x1
		x0 = x1
		x1 = x2
		return x2
	}
}
// 闭包的错误用法
func ClousreTrap() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func ClosureSucc() {
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(j)
		}(i)
	}
}