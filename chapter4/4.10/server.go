// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-30 15:49
// @description:
package main

import (
	"fmt"
	"syscall"
)

func main() {
	ListenFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	err = syscall.Bind(ListenFD, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}
	if err = syscall.Listen(ListenFD, 100); err != nil {
		panic(err)
	}

	for {
		connectFD, _, err := syscall.Accept(ListenFD)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func() {
			fmt.Println(connectFD)
		}()
	}
}
