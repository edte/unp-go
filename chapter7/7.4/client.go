// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-05-08 17:02
// @description:
package main

import (
	"fmt"
	"syscall"
)

// 使用 getsockopt 获取 mss
// connect 前后 mss 不同

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	n, err := syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_MAXSEG)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect before:", n)

	err = syscall.Connect(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}

	n, err = syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_MAXSEG)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect after: ", n)

	_, err = syscall.Write(fd, []byte("hello world"))
	if err != nil {
		panic(err)
	}
}
