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

	e := newEpoll()
	if err := e.add(c.fd); err != nil {
		panic(err)
	}

	if err := e.add(syscall.Stdin); err != nil {
		panic(err)
	}

	for {
		events, _, err := e.wait(-1)
		if err != nil {
			panic(err)
		}

		for i := range events {
			if events[i].Fd == int32(syscall.Stdin) {
				_, err := fmt.Scanln(&str)
				if err != nil {
					panic(err)
				}
				_, err = c.Write([]byte(str))
				if err != nil {
					panic(err)
				}
				continue
			}

			buf := make([]byte, 100)

			_, err = c.Read(buf)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(buf))
		}
	}
}
