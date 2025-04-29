package main

import (
	"fmt"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	msg := make(chan interface{})

	for i := 0; i < 10; i++ {
		go worker(i, msg)
	}

	go producer(msg)

	select {}
}

func producer(msg chan interface{}) {
	for {
		select {
		case <-time.After(time.Second):
			val := time.Now().Unix()
			fmt.Println("on time,send val is", val)
			msg <- val
		}
	}
}

func worker(num int, msg chan interface{}) {
	//fmt.Println("this worker number is", num)

	for {
		select {
		case val := <-msg:
			fmt.Println("cur worker number:", num, "val is:", val)
		}
	}
}
