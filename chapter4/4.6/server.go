// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-29 08:27
// @description:
package main

import (
	"fmt"
	"syscall"
)

// accept  获取连接客户端地址

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
		f, addr, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println(err)
		}

		inet4 := addr.(*syscall.SockaddrInet4)
		fmt.Println(inet4.Addr)
		fmt.Println(inet4.Port)

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
