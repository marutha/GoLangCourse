package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{123, 456, 789}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[1:3], x[5:]...)
	fmt.Println(x)
}
