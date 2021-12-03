/*
 * @Author: yang
 * @Date: 2021-12-03 07:38:46
 * @LastEditTime: 2021-12-03 18:36:42
 * @LastEditors: yang
 * @Description: 我好帅！
 * @FilePath: \project v\server.go
 * -.-
 */
package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}
func (t *Server) Handler(conn net.Conn) {
	fmt.Println("conn sucess !")
}
func (t *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println()
			continue
		}

		go t.Handler(conn)
	}
}
