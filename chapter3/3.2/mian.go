// @program:     unp-go
// @file:        mian.go
// @author:      edte
// @create:      2022-04-28 21:05
// @description:
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 实现 ip 地址字符串和数值的转换

type Option func(ip *IP) error

type IP struct {
	str  string
	data [4]byte
}

func WithByte(d [4]byte) Option {
	return func(ip *IP) error {
		ip.data = d
		return ip.parseByte()
	}
}

func WithStr(s string) Option {
	return func(ip *IP) error {
		ip.str = s
		return ip.parseStr()
	}
}

func New(opts ...Option) (*IP, error) {
	ip := &IP{}

	for _, opt := range opts {
		if err := opt(ip); err != nil {
			return nil, err
		}
	}

	return ip, nil
}

func (i *IP) parseStr() error {
	str := strings.Split(i.str, ".")

	for j, s := range str {
		a, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		i.data[j] = byte(a)
	}

	return nil
}

func (i *IP) parseByte() error {

	for j := range i.data {
		i.str += strconv.Itoa(int(i.data[j])) + "."
	}
	i.str = i.str[:10]
	return nil
}

func (i *IP) String() string {
	return i.str + " : " + fmt.Sprint(i.data)
}

func main() {
	//ip, err := New(WithStr("127.0.0.1"))
	ip, err := New(WithByte([4]byte{192, 18, 9, 1}))
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)
}
