/*
 * @Author: yang
 * @Date: 2021-12-03 18:29:50
 * @LastEditTime: 2021-12-03 19:22:56
 * @LastEditors: yang
 * @Description: 我好帅！
 * @FilePath: \im_golang_pratice\src\main.go
 * -.-
 */
package main

func main() {
	server := NewServer("", 30088)
	server.Start()
}
