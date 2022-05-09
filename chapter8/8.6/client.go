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

// udp client

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
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

	err = syscall.Sendto(fd, []byte(data), 0, addr)
	if err != nil {
		panic(err)
	}

	//fmt.Println("send to...")

	m := make([]byte, 1000)
	n, _, err := syscall.Recvfrom(fd, m, 0)

	fmt.Println(string(m[:n]))
}
