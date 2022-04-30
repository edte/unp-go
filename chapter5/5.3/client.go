// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-30 19:10
// @description:
package main

import (
	"encoding/binary"
	"encoding/json"
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
	var n Num
	buf := NewBuffer()

	for {
		_, err := fmt.Scanf("%d%d", &n.A, &n.B)
		if err != nil {
			panic(err)
		}

		d, err := json.Marshal(n)
		if err != nil {
			panic(err)
		}

		if err = binary.Write(buf, binary.LittleEndian, []byte(d)); err != nil {
			panic(err)
		}
		if err = binary.Write(buf, binary.LittleEndian, []byte("\n")); err != nil {
			panic(err)
		}

		_, err = c.Write(buf.data)
		if err != nil {
			panic(err)
		}

		b := make([]byte, 100)
		_, err = c.Read(b)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}
}
