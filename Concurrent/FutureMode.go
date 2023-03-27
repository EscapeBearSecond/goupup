package main

import (
	"fmt"
	"time"
)

func PutTea() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "茶叶已放入茶杯"
	}()
	return vegetables
}
func BoilingWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(10 * time.Second)
		water <- "水已经烧开"
	}()
	return water
}

// 未来模式
func main() {
	teaCh := PutTea()
	waterCh := BoilingWater()
	fmt.Println("放茶叶，烧水正在进行中，先休息一会~")
	tea := <-teaCh
	water := <-waterCh
	fmt.Println("准备好了，可以沏茶了", tea, water)
}
