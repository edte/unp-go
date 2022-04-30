// @program:     unp-go
// @file:        client.go
// @author:      edte
// @create:      2022-04-26 22:26
// @description:
package main

import "net"

func main() {
	//s := time.Now().String()
	//fmt.Println(len(s))
	// len:52
	//b := []byte(strconv.Itoa(len(s)) + s)
	//fmt.Println(len(b))

	//fmt.Println(len(strconv.Itoa(len(s)) + s + "\n"))

	_, err := net.Listen("tcp", "127.0.0.1")
	if err != nil {
		panic(err)
	}
}
