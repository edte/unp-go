// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-30 19:10
// @description:
package main

import (
	"fmt"
	"syscall"
)

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
	var s string

	for {
		_, err := fmt.Scanln(&s)
		if err != nil {
			panic(err)
		}
		//fmt.Println(s)

		_, err = c.Write([]byte(s))
		if err != nil {
			panic(err)
		}

		buf := make([]byte, 100)
		_, err = c.Read(buf)
		if err != nil {
			panic(err)
		}
		//reader := bufio.NewReader(c)
		//d, _, _ := reader.ReadLine()
		fmt.Println(string(buf))
	}
}
