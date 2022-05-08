// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-30 21:05
// @description:
package main

import (
	"container/list"
	"fmt"
	"syscall"
)

// select io 多路复用重写 client socket 和 stdio
// 服务端也使用 select 优化
// 使用 select 代替多线程

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

	sets := &syscall.FdSet{
		Bits: [16]int64{},
	}

	l := list.New()

	FdZero(sets)
	FdSet(sets, ListenFD)

	for {
		//fmt.Println("-------")

		_, err := syscall.Select(ListenFD+1, sets, nil, nil, nil)
		if err != nil {
			panic(err)
		}

		//fmt.Println(sets)

		if FdIsSet(sets, ListenFD) {
			//fmt.Println("*")

			connectFD, _, err := syscall.Accept(ListenFD)
			//fmt.Println(connectFD)
			if err != nil {
				fmt.Println(err)
				continue
			}
			l.PushBack(connectFD)

			//FdSize(sets)
		}

		//fmt.Println(l.Len())

		for i := 0; i < l.Len(); i++ {
			back := l.Back()
			l.Remove(l.Back())
			echo(back.Value.(int))
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
