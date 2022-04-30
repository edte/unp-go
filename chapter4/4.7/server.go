// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-30 14:56
// @description:
package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

// 信号+fork 子进程

func main() {
	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		fmt.Println("In Child, start...")
	} else {
		fmt.Println("In Father, start...")
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			fmt.Println("exit")
			os.Exit(0)
		case syscall.SIGHUP:
			if err := fork(); err != nil {
				panic(err)
			}
		}
	}
}

func fork() error {
	args := []string{"child"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}
