// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-30 15:39
// @description:
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"syscall"
)

// 实现回显 echo 服务器
// 并且对数据编码
// client 传送 num1,num2
// server 返回 num1+num2

func main() {
	ListenFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	//syscall.SetNonblock(ListenFD, false)

	err = syscall.Bind(ListenFD, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}
	err = syscall.Listen(ListenFD, 100)
	if err != nil {
		panic(err)
	}

	for {
		connectFD, _, err := syscall.Accept(ListenFD)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go echo(newConn(connectFD))
	}
}

func echo(c *Conn) {
	reader := bufio.NewReader(c)

	for {
		//fmt.Println("*")

		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		//fmt.Println(line)

		var a, b int
		_, err = fmt.Sscanf(string(line), "%d%d", &a, &b)
		//fmt.Println(a, b)
		if err != nil {
			panic(err)
		}

		_, err = c.Write([]byte(strconv.Itoa(a + b)))
		if err != nil {
			panic(err)
		}

	}
}
