// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-26 21:27
// @description:
package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"syscall"
)

type socket struct {
	fd int
}

func (s *socket) Read(p []byte) (n int, err error) {
	return syscall.Read(s.fd, p)
}

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOL_SOCKET, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 5)
	if err != nil {
		panic(err)
	}

	err = syscall.Connect(fd, &syscall.SockaddrInet4{
		Port: 9999,
	})
	if err != nil {
		panic(err)
	}

	s := &socket{
		fd: fd,
	}
	b := bufio.NewReader(s)

	for {
		var length int32
		err := binary.Read(b, binary.BigEndian, &length)
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println(b.Buffered())
		if b.Buffered() < 53 {
			continue
		}
		data := make([]byte, 53)
		_, err = b.Read(data)
		if err != nil {
			continue
		}
		fmt.Print(string(data))
	}
}
