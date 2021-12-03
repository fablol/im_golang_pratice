/*
 * @Author: yang
 * @Date: 2021-12-03 19:26:10
 * @LastEditTime: 2021-12-04 00:19:11
 * @LastEditors: yang
 * @Description: 我好帅！
 * @FilePath: \im_golang_pratice\src\user.go
 * -.-
 */
package main

import (
	"net"
	"strings"
)

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
		t.SendMsg("/to <name> <msg>: 发送私聊消息")
		//t.SendMsg("/group <msg>: 发送群聊消息")
		t.SendMsg("/name <name>: 修改昵称")
	} else if msg == "/list" {
		t.SendMsg("在线用户：")
		t.server.mapLock.Lock()
		for _, user := range t.server.OnlineMap {
			sendMsg := "[" + user.Addr + "]" + user.Name
			t.SendMsg(sendMsg)
		}
		t.server.mapLock.Unlock()
	} else if len(msg) > 6 && msg[:6] == "/name " {
		newName := strings.TrimSpace(msg[6:])

		index := strings.Index(newName, " ")
		if index != -1 {
			newName = newName[:index]
		}

		if newName == "" {
			t.SendMsg("昵称不能为空")
			return
		} else if newName == t.Name {
			t.SendMsg("昵称未修改")
			return
		} else if _, ok := t.server.OnlineMap[newName]; ok {
			t.SendMsg("昵称已存在")
			return
		} else {
			t.server.mapLock.Lock()
			delete(t.server.OnlineMap, t.Name)
			t.server.OnlineMap[newName] = t
			t.server.mapLock.Unlock()

			t.Name = newName
			t.SendMsg("您已改名为：" + newName)
		}

	} else if len(msg) > 4 && msg[:4] == "/to " {
		msg = msg[4:]
		index := strings.Index(msg, " ")
		if index == -1 {
			t.SendMsg("消息格式错误")
			return
		} else {
			name := msg[:index]
			msg = msg[index+1:]
			t.server.mapLock.Lock()
			if user, ok := t.server.OnlineMap[name]; ok {
				user.SendMsg("[私聊]" + t.Name + ": " + msg)
			} else {
				t.SendMsg("用户不存在")
			}
			t.server.mapLock.Unlock()
		}
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
