// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-05-08 22:01
// @description:
package main

import (
	"fmt"
	"syscall"
)

// https://www.zhihu.com/question/42308970

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	if err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.TCP_NODELAY, 1000); err != nil {
		panic(err)
	}

	n, err := syscall.GetsockoptInt(fd, syscall.SOL_SOCKET, syscall.TCP_NODELAY)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	err = syscall.Connect(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}

	s := []byte("0123456789")

	for i := range s {
		//fmt.Println(s[i])

		_, err = syscall.Write(fd, []byte{s[i]})
		if err != nil {
			panic(err)
		}
		//time.Sleep(time.Second)
	}
}
