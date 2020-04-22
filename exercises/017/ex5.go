package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48}
	fmt.Println(x)
	x = append(x, x[3:]...)
	fmt.Println(x)
}
