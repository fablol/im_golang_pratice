/*
 * @Author: yang
 * @Date: 2021-12-03 19:26:10
 * @LastEditTime: 2021-12-03 19:31:32
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
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	go user.ListenMessage()

	return user
}

func (t *User) ListenMessage() {
	for {
		msg := <-t.C

		t.conn.Write([]byte(msg + "\n"))
	}
}
