package main

import "fmt"

func main() {
	x := make([]int, 10, 100)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[9] = 10
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 10)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
