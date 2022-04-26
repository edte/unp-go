// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-26 21:27
// @description:
package main

import (
	"syscall"
	"time"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 5)
	if err != nil {
		panic(err)
	}
	err = syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 9999,
	})
	if err != nil {
		panic(err)
	}

	if err = syscall.Listen(fd, 1000); err != nil {
		panic(err)
	}

	f, _, err := syscall.Accept(fd)
	if err != nil {
		panic(err)
	}

	for {
		// 解决粘包问题：
		// 1. 使用分隔符
		_, err = syscall.Write(f, []byte(time.Now().String()+"\n"))
		if err != nil {
			panic(err)
		}
	}
}
