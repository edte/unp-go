// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-05-08 22:40
// @description:
package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"syscall"
)

// 使用 epoll 实现 TCP、UDP 的 echo 服务器

func main() {
	addr := &syscall.SockaddrInet4{
		Port: 1234,
	}

	// tcp
	tcpFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	err = syscall.SetsockoptInt(tcpFD, syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
	if err != nil {
		panic(err)
	}

	err = syscall.Bind(tcpFD, addr)
	if err != nil {
		panic(err)
	}

	err = syscall.Listen(tcpFD, 1000)
	if err != nil {
		panic(err)
	}

	// udp
	udpFD, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		panic(err)
	}

	err = syscall.Bind(udpFD, addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("tcp:", tcpFD, " udp:", udpFD)

	//

	data := make([]byte, 1000)
	buf := make([]byte, 100)

	e := newEpoll()
	if err := e.add(tcpFD); err != nil {
		panic(err)
	}
	if err := e.add(udpFD); err != nil {
		panic(err)
	}

	for {
		events, _, err := e.wait(-1)
		if err != nil {
			panic(err)
		}
		fmt.Println("wait..")

		for i := range events {
			fmt.Println(events[i].Fd)

			if events[i].Fd == int32(tcpFD) {
				fmt.Println("tcp awake..")

				clientFD, _, err := syscall.Accept(tcpFD)
				if err != nil {
					panic(err)
				}

				if err = e.add(clientFD); err != nil {
					panic(err)
				}

				continue
			}

			if events[i].Fd == int32(udpFD) {
				fmt.Println("udp...")

				n, a, err := syscall.Recvfrom(udpFD, data, 0)
				if err != nil {
					panic(err)
				}

				if err = syscall.Sendto(udpFD, data[:n], 0, a); err != nil {
					panic(err)
				}
				continue
			}

			fmt.Println("tcp...")

			n, err := syscall.Read(int(events[i].Fd), buf)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(buf[:n]))

			_, err = syscall.Write(int(events[i].Fd), buf)
			if err != nil {
				panic(err)
			}
		}
	}
}
