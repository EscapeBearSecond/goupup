package channel

import (
	"fmt"
	"time"
)

func NoticeAndSelect() {
	var intChan = make(chan int)
	var strChan = make(chan string)
	var doChan = make(chan struct{})	

	go NoticeAndSelectF1(intChan)
	go NoticeAndSelectF2(strChan)
	go NoticeAndSelectF3(intChan, strChan, doChan)
	time.Sleep(5 * time.Second)
	close(doChan)
}

func NoticeAndSelectF1(intChan chan<- int) {
	for i := 0; i < 100; i++ {
		intChan <- i
	}
}

func NoticeAndSelectF2(strChan chan string) {
	for i := 0; i < 100; i++ {
		strChan <- fmt.Sprintf("this is a string and number is %d", i)
	}
}

func NoticeAndSelectF3(intChan <-chan int, strChan <-chan string, doChan <-chan struct{}) {
	i := 0
	for {
		select {
		case data := <-intChan:
			fmt.Printf("number: %d\n", data)
		case data := <-strChan:
			fmt.Println(data)
		case <-doChan:
			fmt.Println("receive a over signal")
			return
		}
		i++
		fmt.Printf("Total number of times: %d\n", i)
	}
}
