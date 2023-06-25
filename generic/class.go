package generic

import "fmt"

type user struct {
	Name string
	Age  byte
	ID   int64
}

type address struct {
	ID       int
	Province string
	City     string
}

func ClassCase() {
	userMp := make(map[int64]user, 0)
	userMp[1] = user{
		Name: "nick",
		Age:  19,
		ID:   1,
	}
	userMp[2] = user{
		Name: "lucy",
		Age:  20,
		ID:   2,
	}
	userList := mpToList[int64, user](userMp)
	ch := make(chan user)
	go ListPrintln(ch)
	for _, u := range userList {
		ch <- u
	}

	addMp := make(map[int64]address, 0)
	addMp[1] = address{
		ID:       1,
		Province: "广西",
		City:     "桂林",
	}
	addMp[2] = address{
		ID:       2,
		Province: "江苏",
		City:     "连云港",
	}
	addList := mpToList[int64, address](addMp)
	ch1 := make(chan address)
	go ListPrintln(ch1)
	for _, data := range addList {
		ch1 <- data
	}
}
func mpToList[k comparable, T any](mp map[k]T) []T {
	list := make([]T, len(mp))
	var i = 0
	for _, data := range mp {
		list[i] = data
		i++
	}
	return list
}
func mpToList2[k comparable, T any](genericMap GenericMap[k, T]) []T {
	list := make([]T, len(genericMap))
	var i = 0
	for _, data := range genericMap {
		list[i] = data
		i++
	}
	return list
}

func ListPrintln[T any](ch chan T) {
	for data := range ch {
		fmt.Println(data)
	}
}

type GenericMap[k comparable, T any] map[k]T
