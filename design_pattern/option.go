package main

import (
	"fmt"
	"time"
)

// 函数选项模式

type Option func(server *Server)

func WithAddress(address string) Option {
	return func(server *Server) {
		server.Address = address
	}
}
func WithPort(port int) Option {
	return func(server *Server) {
		server.Port = port
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.ReadTimeout = timeout
	}
}
func WithWriteTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.WriteTimeout = timeout
	}
}
func WithTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

func NewServer2(options ...Option) *Server {
	server := &Server{
		Address:      "127.0.0.1",
		Port:         8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		Timeout:      30 * time.Second,
	}
	for _, option := range options {
		option(server)
	}
	return server
}

type Server struct {
	Address      string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Timeout      time.Duration
}

func NewServer(address string, port int, readTimeout, writeTimeout, timeout time.Duration) Server {
	return Server{
		Address:      address,
		Port:         port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Timeout:      timeout,
	}
}
func main() {
	// server := NewServer("127.0.0.1", 8080, 10*time.Second, 20*time.Second, 50*time.Second)
	server := NewServer2(WithAddress("192.168.1.163"), WithPort(9090), WithTimeout(50*time.Second))
	fmt.Println(server)
}
