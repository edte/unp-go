// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-26 21:27
// @description:
package main

import (
	"bufio"
	"io"
	"log"
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
		data, err := b.ReadSlice('\n')
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			} else {
				break
			}
		}
		log.Println("received msg", len(data), "bytes:", string(data))
	}
}
