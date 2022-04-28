// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-28 21:10
// @description:
package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	err = syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}
	if err = syscall.Listen(fd, 1000); err != nil {
		panic(err)
	}

	for {
		_, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println(err)
		}
	}
}
