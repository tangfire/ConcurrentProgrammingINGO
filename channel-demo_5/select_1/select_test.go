package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	default:
		fmt.Println(ctx.Err())
	}

}

func TestChannel(t *testing.T) {
	c := make(chan int)
	notify := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		val := <-c
		fmt.Println("val:", val)

		notify <- struct{}{}
	}()
	isDone := false
	for {
		if isDone {
			break
		}

		select {
		case c <- 1024:
		case <-notify:
			fmt.Println("business is done")
			isDone = true
		case <-time.After(3 * time.Second):
			fmt.Println("business is timeout")
			isDone = true
		}
	}

}

func TestSelect2(t *testing.T) {
	c1 := make(chan int, 1)
	c2 := make(chan int, 2)
	c3 := make(chan int, 3)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		select {
		case v1 := <-c1:
			fmt.Println("v1:", v1)
		case v2 := <-c2:
			fmt.Println("v2:", v2)
		case v3 := <-c3:
			fmt.Println("v3:", v3)
		}

		defer wg.Done()
	}()

	select {
	case c1 <- 1024:
	case c2 <- 1023:
	case c3 <- 1022:
	}

	wg.Wait()

}
