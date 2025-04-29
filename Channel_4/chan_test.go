package main

import (
	"fmt"
	"testing"
)

func TestUnbufferChan(t *testing.T) {
	c := make(chan int) // unbuffer chan

	go func() {
		val := <-c
		fmt.Println(val)
	}()

	c <- 1024

}

func TestBufferChan(t *testing.T) {
	c := make(chan int, 10)

	c <- 1024
	c <- 1023
	c <- 1022

	for v := range c {
		fmt.Println("val:", v)
	}
}

func TestRecChan(t *testing.T) {
	c := make(chan int, 1)
	c <- 1024
	chanTest(c)
}

func chanTest(c <-chan int) {
	val := <-c
	fmt.Println(val)
}

func TestCloseChan(t *testing.T) {
	c := make(chan int, 1)

	c <- 1024

	close(c)

	val1 := <-c
	val2 := <-c
	fmt.Println(val1, val2) // 1024 0

}

func TestCloseChan2(t *testing.T) {
	c := make(chan int, 1)

	c <- 1024

	close(c)

	c <- 2023 // panic: send on closed channel [recovered]

}
