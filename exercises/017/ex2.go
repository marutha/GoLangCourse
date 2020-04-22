package main

import "fmt"

func main() {
	x := []int{10, 11, 12, 13, 14, 50, 51, 52, 53}

	for _, v := range x {
		fmt.Println(v)
	}
	y := x[1:5]
	fmt.Println(y)
	fmt.Println(x[:5])
	fmt.Println(x[5:])
}
