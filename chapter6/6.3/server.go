// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-30 21:05
// @description:
package main

import (
	"fmt"
	"syscall"
)

// epoll io 多路复用重写 client socket 和 stdio

func main() {
	ListenFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}

	syscall.SetNonblock(ListenFD, false)

	syscall.SetsockoptInt(ListenFD, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 99)

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

	e := newEpoll()
	if err := e.add(ListenFD); err != nil {
		panic(err)
	}

	for {
		events, _, err := e.wait(-1)
		if err != nil {
			panic(err)
		}

		for i := range events {
			if events[i].Fd == int32(ListenFD) {
				clientFD, _, err := syscall.Accept(ListenFD)
				if err != nil {
					panic(err)
				}

				if err = e.add(clientFD); err != nil {
					panic(err)
				}

				continue
			}

			echo(int(events[i].Fd))

		}

	}
}

func echo(fd int) {
	buf := make([]byte, 100)
	_, err := syscall.Read(fd, buf)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))

	_, err = syscall.Write(fd, buf)
	if err != nil {
		panic(err)
	}

	buf = make([]byte, 100)

}
