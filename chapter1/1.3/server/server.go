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
	// panic: address family not supported by protocol
	// 协议不存在
	fd, err := syscall.Socket(9999, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 5)
	if err != nil {
		panic(err)
	}
	err = syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 1243,
	})
	if err != nil {
		panic(err)
	}

	if err = syscall.Listen(fd, 1000); err != nil {
		panic(err)
	}

	for {
		f, _, err := syscall.Accept(fd)
		if err != nil {
			panic(err)
		}
		_, err = syscall.Write(f, []byte(time.Now().String()))
		if err != nil {
			panic(err)
		}
		if err := syscall.Close(f); err != nil {
			panic(err)
		}
		break
	}
}
