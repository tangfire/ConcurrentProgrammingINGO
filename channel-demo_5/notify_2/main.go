package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doClean(closed chan struct{}) {
	time.Sleep(2 * time.Second)

	close(closed)
	// 资源释放
	fmt.Println("doClean")
}

func main() {
	fmt.Println("this is main")

	var closing = make(chan struct{})
	var closed = make(chan struct{})

	go func() {
		for {
			select {
			case <-closing:
				return
			default:
				// 模拟业务处理
				time.Sleep(60 * time.Second)
			}
		}
	}()

	notifyC := make(chan os.Signal)

	signal.Notify(notifyC, syscall.SIGINT, syscall.SIGTERM)

	s1 := <-notifyC

	fmt.Println("s1:", s1)

	close(closing)

	go doClean(closed)

	select {
	case <-closed:
	case <-time.After(10 * time.Second):
		fmt.Println("timeout")
	}

	fmt.Println("exit")

}
