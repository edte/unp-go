// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-30 21:05
// @description:
package main

import (
	"fmt"
	"syscall"
)

// select io 多路复用重写 client socket 和 stdio

func main() {
	ListenFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	syscall.SetNonblock(ListenFD, false)

	err = syscall.Bind(ListenFD, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}
	err = syscall.Listen(ListenFD, 100)
	if err != nil {
		panic(err)
	}

	for {
		connectFD, _, err := syscall.Accept(ListenFD)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go echo(connectFD)
	}
}

func echo(fd int) {
	buf := make([]byte, 100)
	for {
		_, err := syscall.Read(fd, buf)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(buf))

		_, err = syscall.Write(fd, buf)
		if err != nil {
			panic(err)
		}

		buf = make([]byte, 100)
	}

}
