package main

import "fmt"

func main() {
	x := []int{10, 11, 12, 13, 14, 50, 51, 52, 53}

	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", x)
}
