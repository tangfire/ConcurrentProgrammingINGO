package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Locker struct {
	c chan struct{}
}

func NewLocker() *Locker {
	lc := Locker{
		c: make(chan struct{}, 1),
	}

	lc.c <- struct{}{}

	return &lc
}

func (lc *Locker) Lock() {
	<-lc.c
}

func (lc *Locker) Unlock() {
	select {
	case lc.c <- struct{}{}:
	default:
		panic("unlock fail")
	}
}

func (lc *Locker) LockTimeout(t time.Duration) bool {
	timer := time.NewTimer(t)
	select {
	case <-lc.c:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}

func TestMutex01(t *testing.T) {
	locker := make(chan struct{})

	a := 100

	for i := 0; i < 100; i++ {
		go func(val *int) {
			<-locker
			*val++
			locker <- struct{}{}
		}(&a)
	}

	locker <- struct{}{}

	time.Sleep(time.Second)

	fmt.Println(a)

}

func TestMutex02(t *testing.T) {
	locker := NewLocker()

	a := 100

	for i := 0; i < 100; i++ {
		go func(val *int) {
			locker.Lock()
			*val++
			locker.Unlock()
		}(&a)
	}

	time.Sleep(time.Second)

	fmt.Println(a)

}

func TestMutex03(t *testing.T) {
	locker := NewLocker()

	locker.Lock()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {

		defer wg.Done()
		isLocked := locker.LockTimeout(2 * time.Second)
		if isLocked {
			fmt.Println("locked")
		} else {
			fmt.Println("lock fail")
		}
	}()

	time.Sleep(time.Second * 3)

	locker.Unlock()

	wg.Wait()
}
