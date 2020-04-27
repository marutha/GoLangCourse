package main

import "fmt"

func main() {
	x := foo()
	y := x()
	fmt.Println(y)
}

func foo() func() int {
	return func() int {
		fmt.Println("Hello")
		return 456
	}
}
