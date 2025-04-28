package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCancelCtx(t *testing.T) {
	ctx := context.Background()

	ctx1, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(5 * time.Second)
		cancel()

	}()

	select {
	case <-ctx1.Done():
		fmt.Println("ctx1 is done")
	}

}

func TestTimeoutCtx(t *testing.T) {
	ctx := context.Background()
	ctx1, cancel := context.WithTimeout(ctx, 5*time.Second)

	// 5s后取消
	defer cancel()

	// 立即取消
	//cancel()

	select {
	case <-ctx1.Done():
		fmt.Println("ctx1 is done")
	}

}

func TestDeadlineCtx(t *testing.T) {
	ctx := context.Background()

	ctx1, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))

	defer cancel()

	select {
	case <-ctx1.Done():
		fmt.Println("ctx1 is done")
	}
}

func TestSomethingCtx(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	doSomething(ctx)

}

func doSomething(ctx context.Context) {
	c := make(chan struct{})

	go func() {
		// do something
		time.Sleep(10 * time.Second)
		c <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		if err == context.DeadlineExceeded {
			fmt.Println("timeout")
		} else {
			fmt.Println("cancel")
		}
		fmt.Println("ctx is done")
	case <-c:
		fmt.Println("do something finish")

	}
}

func TestTimeoutCtx01(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	ctx1 := context.WithValue(ctx, "ctx1", ctx)
	defer cancel()
	select {
	case <-ctx1.Done():
		fmt.Println("ctx1 is done")
	}
}
