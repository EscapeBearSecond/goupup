package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Millisecond * 10) // 模拟慢响应10ms
		fmt.Fprint(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("listen port error, err: %v\n", err)
	}
}
