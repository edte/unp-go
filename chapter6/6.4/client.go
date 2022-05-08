// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-30 21:27
// @description:
package main

import (
	"fmt"
	"syscall"
)

// https://github.com/mindreframer/golang-stuff/blob/master/github.com/pebbe/zmq2/examples/udpping1.go

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
	if err != nil {
		panic(err)
	}
	cli(newConn(fd))
}

type conn struct {
	fd int
}

func newConn(fd int) *conn {
	return &conn{fd: fd}
}

func (c *conn) Write(p []byte) (n int, err error) {
	return syscall.Write(c.fd, p)
}

func (c *conn) Read(p []byte) (n int, err error) {
	return syscall.Read(c.fd, p)
}

func cli(c *conn) {
	var str string
	sets := &syscall.FdSet{
		Bits: [16]int64{},
	}

	for {
		FdZero(sets)
		FdSet(sets, syscall.Stdin)
		FdSet(sets, c.fd)

		_, err := syscall.Select(max(syscall.Stdin, c.fd)+1, sets, nil, nil, nil)
		if err != nil {
			panic(err)
		}

		if FdIsSet(sets, syscall.Stdin) {
			//fmt.Println("stdin is set")

			_, err := fmt.Scanln(&str)
			if err != nil {
				panic(err)
			}

			//fmt.Println(str)

			//fmt.Println("write..")

			_, err = c.Write([]byte(str))
			if err != nil {
				panic(err)
			}
		}

		if FdIsSet(sets, c.fd) {
			//fmt.Println("socket fd is set")

			buf := make([]byte, 100)
			//fmt.Println("send...")
			_, err = c.Read(buf)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(buf))
		}
	}
}
