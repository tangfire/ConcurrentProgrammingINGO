package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var counter int64 = 0
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)

			//fmt.Printf("goroutine num %d incr finished\n", i)
		}(i)
	}

	wg.Wait()

	fmt.Printf("counter %d\n", counter)

	ok := atomic.CompareAndSwapInt64(&counter, 10000, 200)

	if ok {
		fmt.Printf("swap counter %d\n", counter)

	}

}
