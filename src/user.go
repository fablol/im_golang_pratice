/*
 * @Author: yang
 * @Date: 2021-12-03 19:26:10
 * @LastEditTime: 2021-12-03 23:34:12
 * @LastEditors: yang
 * @Description: 我好帅！
 * @FilePath: \im_golang_pratice\src\user.go
 * -.-
 */
package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
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

	go user.ListenMessage()

	return user
}

func (t *User) Online() {
	t.server.mapLock.Lock()
	t.server.OnlineMap[t.Name] = t
	t.server.mapLock.Unlock()

	t.SendMsg("welcome!\n/help: 查看帮助")
	t.server.BroadCast(t, "online")
}

func (t *User) Offline() {
	t.server.mapLock.Lock()
	delete(t.server.OnlineMap, t.Name)
	t.server.mapLock.Unlock()

	t.server.BroadCast(t, "offline")
}

func (t *User) SendMsg(msg string) {
	t.conn.Write([]byte(msg + "\n"))
}

func (t *User) DealMessage(msg string) {
	if msg == "/exit" {
		t.SendMsg("Bye!")
		t.conn.Close()
		//t.Offline()
		return
	} else if msg == "/help" {
		t.SendMsg("/help: 查看帮助")
		t.SendMsg("/list: 查看在线用户")
		t.SendMsg("/exit: 退出聊天室")
	} else if msg == "/list" {
		t.SendMsg("在线用户：")
		t.server.mapLock.Lock()
		for _, user := range t.server.OnlineMap {
			sendMsg := "[" + user.Addr + "]" + user.Name
			t.SendMsg(sendMsg)
		}
		t.server.mapLock.Unlock()
	} else {
		t.server.BroadCast(t, msg)
	}
}

func (t *User) ListenMessage() {
	for {
		msg := <-t.C

		t.conn.Write([]byte(msg + "\n"))
	}
}
