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

// bind 不指定 ip 和端口，则会自动选择 ip 和端口
// 也可以不用 bind，同理
// 要获取 bind 的对应端口，可以使用 getsockname 来获取

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}
	err = syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 0,
	})
	if err != nil {
		panic(err)
	}
	if err := syscall.Listen(fd, 1000); err != nil {
		panic(err)
	}

	name, err := syscall.Getsockname(fd)
	inet4 := name.(*syscall.SockaddrInet4)
	fmt.Println(inet4.Port)
	fmt.Println(inet4.Addr)

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
