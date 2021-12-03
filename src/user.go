/*
 * @Author: yang
 * @Date: 2021-12-03 19:26:10
 * @LastEditTime: 2021-12-03 23:10:48
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

	t.server.BroadCast(t, "online")
}

func (t *User) Offline() {
	t.server.mapLock.Lock()
	delete(t.server.OnlineMap, t.Name)
	t.server.mapLock.Unlock()

	t.server.BroadCast(t, "offline")
}

func (t *User) DoMessage(msg string) {
	t.server.BroadCast(t, msg)
}

func (t *User) ListenMessage() {
	for {
		msg := <-t.C

		t.conn.Write([]byte(msg + "\n"))
	}
}
