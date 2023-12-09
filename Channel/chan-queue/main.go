package main

import "time"

// 使用通道实现一个消息队列
func main() {
	eventChan := make(chan int, 10) // 队列的长度是10
	go produceEvent(eventChan)
	go consumeEvent(eventChan)
	time.Sleep(time.Second * 100) // 等待100s后退出程序
}
func consumeEvent(eventChan chan int) {
	for {
		event := <-eventChan
		println("consume event:", event)
		time.Sleep(time.Second * 2) // 每个消息延迟1s输出
	}
}

func produceEvent(eventChan chan int) {
	i := 0
	for {
		eventChan <- i
		i++
		time.Sleep(time.Second * 1)
	}
}
