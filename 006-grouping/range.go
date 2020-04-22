package main

import "fmt"

func main() {
	x := []int{1, 2, 5, 6, 76}
	fmt.Println(x)
	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}
}
