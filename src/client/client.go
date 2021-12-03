/*
 * @Author: yang
 * @Date: 2021-12-04 00:26:33
 * @LastEditTime: 2021-12-04 01:26:07
 * @LastEditors: yang
 * @Description: 我好帅！
 * @FilePath: \im_golang_pratice\src\client\client.go
 * -.-
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
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

func (c *Client) menu() bool {
	// todo
	return true
}

func (c *Client) Run() {
	c.menu()
	fmt.Println("请输入消息:")
	var msg string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg = scanner.Text()
			break
		}

		if msg == "exit" {
			break
		}
		c.Send(msg)
		msg = ""
	}
}

func (c *Client) Send(msg string) {
	_, err := c.conn.Write([]byte(msg + "\n"))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return
	}

}

func (c *Client) Recv() {
	io.Copy(os.Stdout, c.conn) // 阻塞接收服务端发送的消息
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
	go client.Recv()
	client.Run()
}
