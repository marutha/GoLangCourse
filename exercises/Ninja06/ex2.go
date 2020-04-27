package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	sum := foo(x...)
	fmt.Println(sum)
	fmt.Println(bar(x))
}

func foo(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
