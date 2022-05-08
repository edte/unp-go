// @program:     unp-go
// @file:        item.go
// @author:      edte
// @create:      2022-05-07 22:17
// @description:
package main

import (
	"fmt"
	"syscall"
)

func FdSet(p *syscall.FdSet, i int) {
	p.Bits[i/64] |= 1 << uint(i) % 64
}

func FdIsSet(p *syscall.FdSet, i int) bool {
	return (p.Bits[i/64] & (1 << uint(i) % 64)) != 0
}

func FdZero(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}

func FdSize(p *syscall.FdSet) int {
	for _, bit := range p.Bits {
		fmt.Print(bit, " ")
	}
	fmt.Println()

	//for i := range p.Bits {
	//	p.Bits[i] = 0
	//}

	return 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
