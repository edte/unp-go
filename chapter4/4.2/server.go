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

// 测试 backlog 设置队列长度（半+全）

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
	if err := syscall.Listen(fd, 1); err != nil {
		panic(err)
	}

	name, err := syscall.Getsockname(fd)
	inet4 := name.(*syscall.SockaddrInet4)
	fmt.Println(inet4.Port)
	fmt.Println(inet4.Addr)

	for {
		_, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println(err)
		}

		//name, err := syscall.Getpeername(fd)
		//inet4 := name.(*syscall.SockaddrInet4)
		//fmt.Println(inet4.Port)
		//fmt.Println(inet4.Addr)
		fmt.Println("*")

		//go func() {
		//	data := make([]byte, 1000)
		//	for {
		//		n, err := syscall.Read(f, data)
		//		if err != nil {
		//			fmt.Println(err)
		//		}
		//		fmt.Printf("recived %d bytes\n", n)
		//		fmt.Println(string(data))
		//	}
		//}()
	}
}
