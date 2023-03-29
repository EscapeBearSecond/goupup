package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// 发布订阅模式
func main() {
	p := Publisher{timeout: 100 * time.Millisecond, buffer: 10}
	defer p.Close()
	// 订阅者订阅所有消息
	all := p.Subscribe()
	// 订阅者仅订阅包含golang的消息
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})
	// 发布消息
	p.Publish("hello, world")
	p.Publish("hello, golang!")

	// 加锁
	var wg sync.WaitGroup
	wg.Add(2)

	// 开启goroutine
	go func() {
		for msg := range all {
			_, ok := msg.(string)
			fmt.Println(ok)
		}
		wg.Done()
	}()
	// 开启goroutine
	go func() {
		for msg := range golang {
			v, ok := msg.(string)
			fmt.Println(v)
			fmt.Println(ok)
		}
		wg.Done()
	}()
	p.Close()
	wg.Wait()
}

// Publisher 发布者
type Publisher struct {
	subscribers map[Subscriber]TopicFunc // 订阅者
	buffer      int                      // 订阅者的缓冲区长度
	timeout     time.Duration            // publisher 发送消息的超时时间
	// m 用来保护subscribers
	// 当修改subscribers时（新加订阅或删除订阅）使用写锁
	// 当向某个订阅者发送消息（向Subscriber channel中写入数据）使用读锁
	m sync.RWMutex
}
type (
	// Subscriber 订阅者通道
	Subscriber chan interface{}
	// TopicFunc 主题函数
	TopicFunc func(v interface{}) bool
)

// Subscribe 发布者订阅方法
func (p *Publisher) Subscribe() Subscriber {
	return p.SubscribeTopic(nil)
}

// SubscribeTopic 发布者订阅主题
func (p *Publisher) SubscribeTopic(topicFunc TopicFunc) Subscriber {
	ch := make(Subscriber, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topicFunc
	p.m.Unlock()
	return ch
}

// Delete 删除某个订阅者
func (p *Publisher) Delete(sub Subscriber) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

// Publish 发布者发布消息
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
}

// sendTopic 发送主题
func (p *Publisher) sendTopic(sub Subscriber, topic TopicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}
