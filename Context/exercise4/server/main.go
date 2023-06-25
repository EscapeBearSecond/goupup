package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10)
		fmt.Fprint(res, "slow response")
		return
	}
	fmt.Fprint(res, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
