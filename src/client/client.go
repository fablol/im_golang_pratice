/*
 * @Author: yang
 * @Date: 2021-12-04 00:26:33
 * @LastEditTime: 2021-12-04 00:47:25
 * @LastEditors: yang
 * @Description: 我好帅！
 * @FilePath: \im_golang_pratice\src\client\client.go
 * -.-
 */
package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int

	Name string
	conn net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))

	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn
	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "server ip")
	flag.IntVar(&serverPort, "port", 30088, "server port")
}

func main() {
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("NewClient error")
		return
	}

	fmt.Println("client:", client.ServerIp, client.ServerPort)

	select {}
}
