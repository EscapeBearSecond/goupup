package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int)

	//time.Sleep(time.Second * 1)
	select {
	case num := <-intChan:
		fmt.Println(num)
	case <-time.After(time.Second * 1):
		fmt.Println("超时")
	}
	intChan <- 10
}
