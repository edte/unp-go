// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-05-08 22:40
// @description:
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 实现 UDP echo 服务器
// 展示 UDP 无流量控制，导致的丢包
// 增大 UDP 缓冲区

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		panic(err)
	}
	addr := &syscall.SockaddrInet4{
		Port: 1234,
	}
	cnt := 0

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT)
		sig := <-ch
		switch sig {
		case syscall.SIGINT:
			fmt.Println(cnt)
			os.Exit(0)
		}
	}()

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, 220*1024)
	if err != nil {
		panic(err)
	}

	err = syscall.Bind(fd, addr)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 1000)

	//time.Sleep(time.Second * 3)

	for {
		_, _, err := syscall.Recvfrom(fd, data, 0)
		if err != nil {
			panic(err)
		}
		cnt++
	}
}
