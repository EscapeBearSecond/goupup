package main

import (
	"log"
	"os"
)

func main() {
	log.SetPrefix("Test:")                                  // 设置日志前缀
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds) // 设置日志格式
	f, err := os.OpenFile("./go.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 777)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)     // 设置日志输出位置
	log.Println("Hello") // 打印并换行
	log.Print("hello")   // 普通打印
	name := "alice"
	log.Printf("hello %s", name) // 格式化打印

	// log.Fatal("something error") // 打印并退出程序
	log.Panic("something panic") // 打印并抛出异常信息
}
