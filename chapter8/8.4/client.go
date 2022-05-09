// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-05-08 22:40
// @description:
package main

import (
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		panic(err)
	}
	addr := &syscall.SockaddrInet4{
		Port: 1234,
	}

	data := make([]byte, 1400)

	for i := 0; i < 200000; i++ {
		err = syscall.Sendto(fd, data, 0, addr)
		if err != nil {
			panic(err)
		}
	}
}
