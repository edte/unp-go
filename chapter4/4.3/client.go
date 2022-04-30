// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-29 08:18
// @description:
package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
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
	//for {
	//	n, err := syscall.Write(fd, []byte("hello world"))
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("send %d bytes\n", n)
	//}
	time.Sleep(time.Hour * 5)
}
