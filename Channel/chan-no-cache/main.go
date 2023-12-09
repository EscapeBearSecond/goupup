package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	value int
	mutex sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool)

	data := Data{value: 0}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go readData(i, &data, ch, &wg)
	}

	wg.Wait()
	close(ch)
}

func readData(id int, data *Data, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		<-ch              // 等待写入许可
		data.mutex.Lock() // 加锁，对资源进行写操作
		data.value++
		fmt.Printf("Goroutine %d writing data: %d\n", id, data.value)
		data.mutex.Unlock() // 解锁

		ch <- true // 释放写入许可

		time.Sleep(time.Millisecond * 500) // 模拟其他处理
	}
}
