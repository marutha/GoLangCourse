package main

import "fmt"

func main() {
	a := 1
	for a < 100 {
		a += 2
		fmt.Println(a)
	}
}
