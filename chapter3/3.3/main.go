// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-28 21:06
// @description:
package main

import (
	"bufio"
	"io"
	syscall "syscall"
)

// 实现 read n 字节，write n 字节，read 一行
// readN,writeN,readLine

// read 和 write 时，直接指定 slice 的长度为 n 即可，或者直接创建一个长度为 n
// 的数组，这样就自动实现读写 n byte 了
func readN(fd int, n int) ([]byte, int, error) {
	data := make([]byte, n)
	n, err := syscall.Read(fd, data)
	return data, n, err
}

func writeN(fd int, data []byte, n int) (int, error) {
	return syscall.Write(fd, data)
}

func readLine(reader io.Reader) ([]byte, bool, error) {
	r := bufio.NewReader(reader)
	return r.ReadLine()
}

func main() {
}
