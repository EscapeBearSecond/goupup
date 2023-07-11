package _sync

import (
	"fmt"
	"sync"
)

func MapCase() {
	mp := sync.Map{}

	fmt.Println(mp.LoadOrStore("name", "jack"))
	fmt.Println(mp.LoadOrStore("name", "jerry"))
	fmt.Println(mp.LoadOrStore("name", "mick"))

}
