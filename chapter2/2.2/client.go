// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-28 21:18
// @description:
package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		panic(err)
	}
	err = syscall.Connect(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}
	for {
		n, err := syscall.Write(fd, []byte("hello "))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("write: %d bytes\n", n)
		}
	}
}
