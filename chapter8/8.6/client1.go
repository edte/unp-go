// @program:     unp-go
// @file:        client1.go
// @author:      edte
// @create:      2022-05-09 21:45
// @description:
package main

import (
	"fmt"
	"syscall"
)

// tcp client

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}
	if err = syscall.Connect(fd, &syscall.SockaddrInet4{Port: 1234}); err != nil {
		panic(err)
	}

	var s string
	_, err = fmt.Scanf("%s", &s)
	if err != nil {
		panic(err)
	}

	//fmt.Println("write...")
	_, err = syscall.Write(fd, []byte(s))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 100)
	n, err := syscall.Read(fd, buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
}
