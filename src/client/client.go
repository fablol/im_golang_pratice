/*
 * @Author: yang
 * @Date: 2021-12-04 00:26:33
 * @LastEditTime: 2021-12-04 00:40:41
 * @LastEditors: yang
 * @Description: 我好帅！
 * @FilePath: \im_golang_pratice\src\client\client.go
 * -.-
 */
package main

import (
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort string

	Name string
	conn net.Conn
}

func NewClient(serverIp, serverPort string) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	conn, err := net.Dial("tcp", serverIp+":"+serverPort)
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn
	return client
}

func main() {
	client := NewClient("127.0.0.1", "30088")
	if client == nil {
		fmt.Println("NewClient error")
		return
	}

	fmt.Println("client:", client.ServerIp, client.ServerPort)

	select {}
}
