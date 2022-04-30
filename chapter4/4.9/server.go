// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-30 15:16
// @description:
package main

import (
	"fmt"
	"syscall"
)

// 多进程并发模型实现

func fork() (pid uintptr) {
	id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	return id
}

func handle(fd int) {
	for {
		data := make([]byte, 1000)
		n, err := syscall.Read(fd, data)
		if err != nil {
			panic(err)
		}
		if n <= 0 {
			continue
		}
		fmt.Println(string(data))
	}
}

func main() {
	listenFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	err = syscall.Bind(listenFD, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}

	if err = syscall.Listen(listenFD, 1000); err != nil {
		panic(err)
	}

	for {
		connectFD, _, err := syscall.Accept(listenFD)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if pid := fork(); pid == 0 {
			syscall.Close(listenFD)
			handle(connectFD)
			syscall.Close(connectFD)
			syscall.Exit(0)
		}
		syscall.Close(connectFD)
	}

	//id := fork()
	//if id == 0 {
	//	fmt.Println("In child,pid: ", id)
	//} else {
	//	fmt.Println("In parent,pid: ", id)
	//}
}
