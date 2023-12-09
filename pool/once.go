package pool

import (
	"fmt"
	"sync"
)

type OnceMap struct {
	sync.Once
	Data map[string]int
}

// LoadData load data
// When multiple threads reach m.Do at the same time,
// only one thread will continue to execute the method in m.Do
func (m *OnceMap) LoadData() {
	m.Do(func() {
		list := []string{"A", "B", "C", "D"}
		for _, item := range list {
			if _, ok := m.Data[item]; !ok {
				m.Data[item] = 0
			}
			m.Data[item] += 1
		}
	})
}
func OnceCase() {
	om := &OnceMap{
		Data: make(map[string]int),
	}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			om.LoadData()
		}()
	}
	wg.Wait()
	fmt.Println(om.Data)
}
