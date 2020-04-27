package main

import "fmt"

func main() {
	g := func(xi []int) int {
		total := 0
		for _, v := range xi {
			total += v
		}
		return total
	}
	result := foo(g)
	fmt.Println(result)
}

func foo(f func(xi []int) int) int {
	i := []int{1, 2, 3, 4}
	n := f(i)
	n++
	return n
}
