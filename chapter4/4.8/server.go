// @program:     unp-go
// @file:        server.go
// @author:      edte
// @create:      2022-04-30 13:16
// @description:
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// 多进程 fork + 信号量实现热重启
// https://www.hitzhangjie.pro/blog/2020-08-28-go%E7%A8%8B%E5%BA%8F%E5%A6%82%E4%BD%95%E5%AE%9E%E7%8E%B0%E7%83%AD%E9%87%8D%E5%90%AF/
// https://cloud.tencent.com/developer/article/1388556
// https://kalifun.github.io/2020/12/28/go-server-ru-he-you-ya-de-qi-dong-re-chong-qi/

var (
	server   *http.Server
	listener net.Listener
	child    = flag.Bool("child", false, "")
)

//func init() {
//	updatePidFile()
//}
//
//func updatePidFile() {
//	sPid := fmt.Sprint(os.Getpid())
//	tmpDir := os.TempDir()
//	if err := procExsit(tmpDir); err != nil {
//		fmt.Printf("pid file exists, update\n")
//	} else {
//		fmt.Printf("pid file NOT exists, create\n")
//	}
//	pidFile, _ := os.Create(tmpDir + "/gracefulRestart.pid")
//	defer pidFile.Close()
//	pidFile.WriteString(sPid)
//}
//
//// 判断进程是否启动
//func procExsit(tmpDir string) (err error) {
//	pidFile, err := os.Open(tmpDir + "/gracefulRestart.pid")
//	defer pidFile.Close()
//	if err != nil {
//		return
//	}
//
//	filePid, err := ioutil.ReadAll(pidFile)
//	if err != nil {
//		return
//	}
//	pidStr := fmt.Sprintf("%s", filePid)
//	pid, _ := strconv.Atoi(pidStr)
//	if _, err := os.FindProcess(pid); err != nil {
//		fmt.Printf("Failed to find process: %v\n", err)
//		return
//	}
//
//	return
//}

func main() {
	flag.Parse()

	// 启动监听
	http.HandleFunc("/hello", HelloHandler)
	server = &http.Server{Addr: ":8081"}

	var err error
	if *child {
		fmt.Println("In Child, Listening...")

		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		fmt.Println("In Father, Listening...")

		listener, err = net.Listen("tcp", server.Addr)
	}
	if err != nil {
		fmt.Printf("Listening failed: %v\n", err)
		return
	}

	// 单独go程启动server
	go func() {
		err = server.Serve(listener)
		if err != nil {
			fmt.Printf("server.Serve failed: %v\n", err)
		}
	}()

	//监听系统信号
	singalHandler()
	fmt.Printf("singalHandler end\n")

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(20 * time.Second)
	//for i := 0; i < 20; i++ {
	//	log.Printf("working %v\n", i)
	//	time.Sleep(1 * time.Second)
	//}
	w.Write([]byte("hello world"))
}

func singalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		sig := <-ch
		fmt.Printf("signal: %v\n", sig)

		ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Printf("stop")
			signal.Stop(ch)
			server.Shutdown(ctx)
			fmt.Printf("graceful shutdown\n")
			return
		case syscall.SIGHUP:
			// reload
			log.Printf("restart")
			err := restart()
			if err != nil {
				fmt.Printf("graceful restart failed: %v\n", err)
			}
			//更新当前pidfile
			//updatePidFile()
			server.Shutdown(ctx)
			fmt.Printf("graceful reload\n")
			return
		}
	}
}

func restart() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return fmt.Errorf("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"-child"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}
