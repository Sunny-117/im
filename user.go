package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server // 当前用户是属于哪一个server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()
	return user
}

func (this *User) Online() {
	// 用户上线，将用户加入到onLineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播当前用户上线消息
	this.server.BroadCast(this, "login success")
}

func (this *User) Offline() {
	// 用户下线，将用户从onLineMap删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	// 广播当前用户上线消息
	this.server.BroadCast(this, "user offline")
}

// send message to current user client
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// query current online users
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := user.Name + " is online..." + "\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else {
		this.server.BroadCast(this, msg)
	}
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
