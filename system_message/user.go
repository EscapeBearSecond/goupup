package main

import "net"

type User struct {
	Name     string
	Addr     string
	Messages chan string
	Conn     net.Conn
}

// 创建用户
func NewUser(conn net.Conn) *User {
	user := &User{
		Name:     conn.RemoteAddr().String(),
		Addr:     conn.RemoteAddr().String(),
		Messages: make(chan string),
		Conn:     conn,
	}
	go user.ListenMessage()
	return user
}

// 监听当前用户的消息
func (u *User) ListenMessage() {
	for {
		// 从管道中得到消息
		msg := <-u.Messages
		// 将消息写入到当前用户的连接
		u.Conn.Write([]byte(msg + "\n"))
	}
}
