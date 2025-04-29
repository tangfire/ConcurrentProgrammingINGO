package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Person struct {
	sync.Mutex
	Name string
	Age  int
}

func TestAgeInc(t *testing.T) {
	p := Person{
		Name: "tangfire",
		Age:  0,
	}

	for i := 0; i < 1000; i++ {
		go func() {
			defer p.Unlock()
			p.Lock()
			p.Age++
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("age:%d\n", p.Age)
}

func TestMutex(t *testing.T) {
	var mutex sync.Mutex

	var counter int = 0

	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			counter++
			defer mutex.Unlock()
		}()

	}

	time.Sleep(time.Second)

	fmt.Printf("counter = %d\n", counter)
}
