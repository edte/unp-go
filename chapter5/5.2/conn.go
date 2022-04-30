// @program:     unp-go
// @file:        conn.go
// @author:      edte
// @create:      2022-04-30 20:28
// @description:
package main

import "syscall"

type Conn struct {
	fd int
}

func newConn(fd int) *Conn {
	return &Conn{fd: fd}
}

func (c *Conn) Write(p []byte) (n int, err error) {
	return syscall.Write(c.fd, p)
}

func (c *Conn) Read(p []byte) (n int, err error) {
	return syscall.Read(c.fd, p)
}
