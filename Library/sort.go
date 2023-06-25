package Library

import (
	"fmt"
	"sort"
)

type user struct {
	ID   int64
	Name string
	Age  int
}
type sortUserByAge []user

func (userList sortUserByAge) Len() int {
	return len(userList)
}
func (userList sortUserByAge) Swap(i, j int) {
	userList[i], userList[j] = userList[j], userList[i]
}
func (userList sortUserByAge) Less(i, j int) bool {
	return userList[i].Age < userList[j].Age
}
func SortCase() {

	userList := []user{
		{ID: 1, Name: "Dyg", Age: 21},
		{ID: 8, Name: "Dyg", Age: 26},
		{ID: 2, Name: "Dyg", Age: 29},
		{ID: 3, Name: "Dyg", Age: 22},
		{ID: 7, Name: "Dyg", Age: 24},
		{ID: 6, Name: "Dyg", Age: 27},
		{ID: 4, Name: "Dyg", Age: 20},
	}
	sort.Slice(userList, func(i, j int) bool {
		return userList[i].ID < userList[j].ID
	})
	fmt.Println(userList)
	sort.Sort(sortUserByAge(userList))
	fmt.Println(userList)
}
