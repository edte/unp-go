// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-26 21:27
// @description:
package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOL_SOCKET, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 5)
	if err != nil {
		panic(err)
	}

	err = syscall.Connect(fd, &syscall.SockaddrInet4{
		Port: 1243,
	})
	if err != nil {
		panic(err)
	}

	data := make([]byte, 1000)

	for {
		n, err := syscall.Read(fd, data)
		if n < 0 {
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}
}
