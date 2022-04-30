// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-30 15:48
// @description:
package main

import (
	"fmt"
	"syscall"
)

// 查看 client 不 bind，查看本地端口范围
// 我多次查看，大概是 559xx 的样子

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}
	err = syscall.Connect(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}
	addr, err := syscall.Getsockname(fd)
	if err != nil {
		panic(err)
	}
	inet4 := addr.(*syscall.SockaddrInet4)
	fmt.Println(inet4.Port)
}
