package main

import "fmt"

var a, b int

func func1() {
	a = 1
	b = 2
}

func func2() {
	print(b)
	print(a)
}

func main() {
	go func1()
	func2()
	fmt.Println("")
}
