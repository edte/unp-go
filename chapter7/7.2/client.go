// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-05-08 21:34
// @description:
package main

import "syscall"

// 默认

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	err = syscall.SetsockoptLinger(fd, syscall.SOL_SOCKET, syscall.SO_LINGER, &syscall.Linger{
		Onoff:  0,
		Linger: 0,
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

	_, err = syscall.Write(fd, []byte("hello world"))
	if err != nil {
		panic(err)
	}
}
