package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRWMutes(t *testing.T) {
	var counter int = 0

	var rwmutex sync.RWMutex

	for i := 0; i < 1000; i++ {
		go func() {
			rwmutex.Lock()
			defer rwmutex.Unlock()
			counter++
		}()
	}

	time.Sleep(time.Second)

	fmt.Printf("Counter: %d\n", counter)
}
