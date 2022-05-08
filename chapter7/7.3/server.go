// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-05-08 22:01
// @description:
package main

import (
	"fmt"
	"syscall"
)

// TCP_NODELAY 选项， Nagle 算法体验
func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}
	err = syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}
	if err := syscall.Listen(fd, 1000); err != nil {
		panic(err)
	}

	for {
		f, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println(err)
		}

		go func() {
			data := make([]byte, 1000)
			n, err := syscall.Read(f, data)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("recived %d bytes\n", n)
			fmt.Println(string(data))
		}()
	}
}
