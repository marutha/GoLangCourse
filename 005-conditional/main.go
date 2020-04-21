package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}

	if 2 == 2 {
		fmt.Println("003")
	}

	if 2 != 2 {
		fmt.Println("004")
	}

	if !true {
		fmt.Println("005")
	}

	if !false {
		fmt.Println("006")
	}
}
