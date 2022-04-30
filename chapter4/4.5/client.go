// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-29 08:18
// @description:
package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	i := 0

	for {
		err = syscall.Bind(fd, &syscall.SockaddrInet4{
			Port: i,
		})
		if err != nil {
			panic(err)
		}

		err = syscall.Connect(fd, &syscall.SockaddrInet4{
			Port: 1234,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("connect succeed")
		i++

		//for {
		//	n, err := syscall.Write(fd, []byte("hello world"))
		//	if err != nil {
		//		panic(err)
		//	}
		//	fmt.Printf("send %d bytes\n", n)
		//}

	}

}
