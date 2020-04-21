package main

import "fmt"

func main() {
	x := 42

	if x > 42 {
		fmt.Println("x is greater")
	} else if x < 42 {
		fmt.Println("x is lesser")
	} else {
		fmt.Println("x is equal")
	}
}
