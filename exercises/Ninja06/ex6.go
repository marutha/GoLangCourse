package main

import "fmt"

func main() {
	x := func() int {
		return 456
	}
	fmt.Println(x())
}
