// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-30 21:27
// @description:
package main

import "syscall"

func main() {
	syscall.Shutdown(0, syscall.SHUT_WR)
}
