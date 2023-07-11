package _sync

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	count int64
}
// 累加操作
func (counter *atomicCounter) Inc() {
	atomic.AddInt64(&counter.count, 1)
}
func (counter *atomicCounter) Load() int64 {
	return atomic.LoadInt64(&counter.count)
}

// 计数器，统计协程数量
func AtomicCase() {
	var count = atomicCounter{}
	// var locker = sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// locker.Lock()
			count.Inc()
			// locker.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count.Load())
}

// 
func AtomicCase1() {
	list := []string{"A", "B", "C", "D"}
	// 定义一个原子值
	var atomicMp atomic.Value
	// locker := sync.Mutex{}
	// 定义一个集合
	mp := make(map[string]int, 0)
	// 将集合存储到原子值
	atomicMp.Store(&mp)
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			FailSwap:
			m := atomicMp.Load().(*map[string]int)
			m1 := map[string]int{}
			for k, v := range *m {
				m1[k] = v
			}
			// locker.Lock()
			for _, v := range list {
				if _, ok := m1[v]; !ok {
					m1[v] = 0
				}
				m1[v] += 1
			}
			// locker.Unlock()
			succSwap := atomicMp.CompareAndSwap(m, &m1)
			if !succSwap {
				// 没有替换成功，重新执行替换逻辑
				goto FailSwap
			}
			
		}()
	}
	wg.Wait()
	fmt.Println(atomicMp.Load())
}