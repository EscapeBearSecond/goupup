package _defer

import (
	"fmt"
	"log"
	"os"
)

// 参数预定义
func DeferCase1() {
	var i int = 100
	defer func() {
		fmt.Println("defer1", i)
	}()
	defer func(j int) {
		fmt.Println("defer 2", j)
	}(i)
	i = 200
	fmt.Println("main", i)
}
func FileCase() {
	file, err := os.Open("Readme.md")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
}



