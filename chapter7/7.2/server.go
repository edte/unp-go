// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-05-08 21:46
// @description:
package main

import (
	"fmt"
	"syscall"
)

// 通过设置 SO_LINGER 选项
// 来选择不同的 close 方式
// 参考 https://breezetemple.github.io/2019/07/04/tcp-option-SO-LINGER/

// 默认状态时，l_onoff == 0，则调用 close 后，会把剩下的数据发送完毕才通过 fin 关闭
// 设置 l_onoff != 0 && l_linger == 0 时，则会立刻发送 RST 关闭，不再发送缓冲区中的数据，并且直接进入 CLOSED 状态，没有 TIME_WAIT 态
// 设置 l_onoff != 0 && l_linger != 0 时，则会在特定时间内尝试发送缓冲区中的数据，规定时间内发送完毕，则按 FIN 关闭，超时则按 RST 关闭

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		panic(err)
	}

	if err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, -1); err != nil {
		panic(err)
	}

	err = syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 1234,
	})
	if err != nil {
		panic(err)
	}
	if err := syscall.Listen(fd, 100); err != nil {
		panic(err)
	}

	for {
		connFD, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println(err)
		}

		data := make([]byte, 1000)
		n, err := syscall.Read(connFD, data)
		if err != nil {
			panic(err)
		}
		if n < 0 {
			fmt.Println("read error")
			syscall.Close(connFD)
		}
		if n == 0 {
			fmt.Println("read EOF")
			syscall.Close(connFD)
		}
		if n > 0 {
			fmt.Println(string(data[:n]))
		}

	}

}
