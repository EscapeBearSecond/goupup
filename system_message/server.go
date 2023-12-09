package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port string
	// 在线人数
	OnlineUsers map[string]*User
	sync.RWMutex
	// 消息广播
	Message chan string
}

func NewServer(ip string, port string) *Server {
	return &Server{
		Ip:          ip,
		Port:        port,
		OnlineUsers: make(map[string]*User),
		Message:     make(chan string, 100),
	}
}

// 监听消息
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message
		s.BroadCast(msg)
	}
}

// 建立连接后的处理
func (s *Server) Handler(conn net.Conn) {
	// fmt.Println("连接建立成功")
	user := NewUser(conn)
	// 用户上线
	s.Lock()
	s.OnlineUsers[user.Name] = user
	s.Unlock()
	// 发送上线消息
	s.ReceiveMessage(user, "上线")
}

// 广播消息
func (s *Server) BroadCast(message string) {
	s.RLock()
	for _, u := range s.OnlineUsers {
		u.Messages <- message
	}
	s.RUnlock()
}

// 接收消息
func (s *Server) ReceiveMessage(user *User, message string) {
	s.Message <- fmt.Sprintf("%s:%s", user.Name, message)
}

// 启动服务器
func (s *Server) Start() {
	// listen
	fmt.Println("开始启动服务器.....")
	lintener, err := net.Listen("tcp", s.Ip+":"+s.Port)
	if err != nil {
		log.Fatalf("listen error, err = %v", err)
	}
	go s.ListenMessage()
	// close listen
	defer lintener.Close()
	for {
		// accept
		conn, err := lintener.Accept()
		if err != nil {
			log.Printf("accecpt error, errr = %v", err)
			continue
		}
		// doHandler
		go s.Handler(conn)
	}

}
