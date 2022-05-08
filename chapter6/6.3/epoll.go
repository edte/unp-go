// @program:     unp-go
// @file:        epoll.go
// @author:      edte
// @create:      2022-05-08 16:02
// @description:
package main

import "syscall"

// https://blog.csdn.net/ALakers/article/details/116849256
// https://www.jianshu.com/p/ee381d365a29
// https://github.com/tidwall/evio

type epoll struct {
	epfd int
}

func newEpoll() *epoll {
	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		panic(err)
	}

	return &epoll{
		epfd: epfd,
	}

}

func (e *epoll) add(fd int) error {
	return e.do(fd, syscall.EPOLL_CTL_ADD)
}

func (e *epoll) del(fd int) error {
	return e.do(fd, syscall.EPOLL_CTL_DEL)
}

func (e *epoll) do(fd int, op int) error {
	err := syscall.EpollCtl(e.epfd, op, fd, &syscall.EpollEvent{
		Events: syscall.EPOLLIN,
		Fd:     int32(fd),
	})
	return err
}

func (e *epoll) wait(msec int) (events []syscall.EpollEvent, n int, err error) {
	data := make([]syscall.EpollEvent, 20)

	n, err = syscall.EpollWait(e.epfd, data, msec)

	return data[:n], n, err
}
