package pool

import (
	"log"
	"math/rand"
	"sync"
)

const (
	ON  = 1
	OFF = 0
)

type Conn struct {
	ID     int64
	Target string
	Status int
}

func NewConn(target string) *Conn {
	conn := &Conn{
		ID:     rand.Int63(),
		Target: target,
		Status: ON,
	}
	log.Println("create a new connection and the connection ID is ", conn.ID)
	return conn
}

type ConnPool struct {
	sync.Pool
}

// GetPool Get connection poll
// The purpose of the New method is to create a new connection when there are no connections in the pool
// New method is required
func GetPool(target string) *ConnPool {
	log.Println("get connection pool...")
	return &ConnPool{
		Pool: sync.Pool{
			New: func() any {
				return NewConn(target)
			},
		},
	}
}

func (p *ConnPool) Get() *Conn {
	conn := p.Pool.Get().(*Conn)
	if conn.Status == OFF {
		conn = p.Pool.New().(*Conn)
	}
	return conn
}
func (p *ConnPool) Put(conn *Conn) {
	if conn.Status == OFF {
		return
	}
	p.Pool.Put(conn)
	log.Println("put the connection to the connection poll successfully and the connection ID is", conn.ID)
}
func PoolCase() {
	target := "127.0.0.1"
	pool := GetPool(target)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				conn := pool.Get()
				log.Println("get the connection successfully and the connection ID is ", conn.ID)
				pool.Put(conn)
			}
		}()
	}
}
