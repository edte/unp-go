// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-05-08 22:40
// @description:
package main

import (
	"fmt"
	"syscall"
)

// 实现 UDP echo 服务器

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		panic(err)
	}
	addr := &syscall.SockaddrInet4{
		Port: 1234,
	}
	err = syscall.Bind(fd, addr)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 1000)

	for {
		n, a, err := syscall.Recvfrom(fd, data, 0)
		if err != nil {
			panic(err)
		}

		fmt.Println(a)

		fmt.Println(string(data[:n]))

		if err = syscall.Sendto(fd, data[:n], 0, a); err != nil {
			panic(err)
		}

		data = make([]byte, 1000)
	}
}
