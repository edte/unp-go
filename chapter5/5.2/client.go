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

func cli(c *Conn) {
	var a, b int

	for {
		_, err := fmt.Scanf("%d%d", &a, &b)
		if err != nil {
			panic(err)
		}

		//fmt.Println(fmt.Sprintf("%d %d\n", a, b))

		_, err = c.Write([]byte(fmt.Sprintf("%d %d\n", a, b)))
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
