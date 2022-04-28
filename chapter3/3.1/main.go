// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-28 21:05
// @description:
package main

import (
	"fmt"
	"unsafe"
)

// 大小端判断

func main() {
	i := int32(0x01020304)
	p := unsafe.Pointer(&i)
	b := (*byte)(p)
	if *b == 01 {
		fmt.Println("small")
	} else if *b == 04 {
		fmt.Println("big")
	}
}
