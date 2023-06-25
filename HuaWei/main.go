package main

import "fmt"

func main() {

	var N int
	fmt.Scanf("%d\n", &N)
	var num = make([][]int, N)
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d\n", &num[i])
	}
}
