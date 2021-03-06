/*
 * @Author: yang
 * @Date: 2021-12-03 07:38:46
 * @LastEditTime: 2021-12-04 01:16:25
 * @LastEditors: yang
 * @Description: 我好帅！
 * @FilePath: \im_golang_pratice\src\server\server.go
 * -.-
 */
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

func (t *Server) ListenMessager() {
	for {
		msg := <-t.Message

		t.mapLock.Lock()
		for _, cli := range t.OnlineMap {
			cli.C <- msg
		}
		t.mapLock.Unlock()
	}
}

func (t *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	t.Message <- sendMsg
}

func (t *Server) Handler(conn net.Conn) {
	fmt.Println("conn sucess !")

	user := NewUser(conn, t)

	user.Online()

	isLive := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)

			if n == 0 {
				user.Offline()
				return
			}

			if err != nil {
				fmt.Println("read err:", err)
				return
			}

			msg := string(buf[:n-1])
			user.DealMessage(msg)

			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:
			// fmt.Println("conn is live")
		case <-time.After(time.Minute * 10):
			user.SendMsg("你已经超时了，请重新登录")

			close(user.C)
			conn.Close()
			return
		}
	}
}

func (t *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	defer listener.Close()

	go t.ListenMessager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println()
			continue
		}

		go t.Handler(conn)
	}
}
