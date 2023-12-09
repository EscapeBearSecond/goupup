package _sync

import (
	"fmt"
	"sync"
	"time"
)

func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{})
	go ReadList(&list, cond)
	go ReadList(&list, cond)
	go ReadList(&list, cond)
	time.Sleep(time.Second * 10)
	// go InitList(&list, cond)
}

func InitList(list *[]int, cond *sync.Cond) {
	// caller can hold locks or not hold locks
	cond.L.Lock()
	defer cond.L.Unlock()
	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	// wake up all waiting concurrent programs
	cond.Broadcast()
	// cond.Signal()
}

func ReadList(list *[]int, cond *sync.Cond) {
	// called must hold locks
	cond.L.Lock()
	defer cond.L.Unlock()
	for len(*list) == 0 {
		fmt.Println("read_list waiting")
		cond.Wait()
	}
	fmt.Println("list data: ", *list)
}

type CondQueue struct {
	list *[]int
	*sync.Cond
}

func (q *CondQueue) GetMany(n int) []int {
	q.L.Lock()
	defer q.L.Unlock()
	for len(*q.list) < n {
		q.Wait()
	}
	list := (*q.list)[:n]
	*q.list = (*q.list)[n:]
	return list
}
func (q *CondQueue) Put(item int) {
	q.L.Lock()
	defer q.L.Unlock()
	*q.list = append(*q.list, item)
	q.Signal()
}
func CondQueueCase() {
	q := &CondQueue{
		list: &[]int{},
		Cond: sync.NewCond(&sync.Mutex{}),
	}
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			list := q.GetMany(n)
			fmt.Printf("%d: %d\n", n, list)
		}(i)
	}

	for i := 0; i < 50; i++ {
		q.Put(i)
		// time.Sleep(time.Second)
	}
	wg.Wait()
}
