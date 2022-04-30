// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-30 15:34
// @description:
package main

import "syscall"

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	err = syscall.Connect(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}

	syscall.Write(fd, []byte("hello world"))

}
