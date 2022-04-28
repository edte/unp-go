// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-28 21:04
// @description:
package main

import (
	"fmt"
	"syscall"
)

// 向 UDP 缓冲区一直写数据

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		panic(err)
	}

	err = syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}

	for {
		data := make([]byte, 10000)
		_, err := syscall.Read(fd, data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data[:]))
	}
}
