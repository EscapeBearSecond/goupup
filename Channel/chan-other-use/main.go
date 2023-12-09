package main

import (
	"fmt"
	"time"
)

func main() {

}

// 计时器机制
func TimerChan() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("Timer expired")
}

// 超时机制
func TimeoutChan() {
	<-time.After(time.Second * 2)
}
