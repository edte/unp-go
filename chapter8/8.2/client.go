// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-05-08 22:40
// @description:
package main

import (
	"fmt"
	"syscall"
)

// 使用 connect，read，write 代替 UDP 的 sendto,recvfrom
// 前者其实是协议无关的

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		panic(err)
	}
	addr := &syscall.SockaddrInet4{
		Port: 1234,
	}

	var data string
	_, err = fmt.Scanf("%s", &data)
	if err != nil {
		panic(err)
	}

	//fmt.Println("send to ...")

	err = syscall.Connect(fd, addr)
	if err != nil {
		panic(err)
	}

	_, err = syscall.Write(fd, []byte(data))
	if err != nil {
		panic(err)
	}

	m := make([]byte, 1000)
	n, err := syscall.Read(fd, m)
	fmt.Println(string(m[:n]))
}
