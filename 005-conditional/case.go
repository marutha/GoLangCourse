package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("This should print")
		fallthrough
	case (4 == 4):
		fmt.Println("This should print")
	}
}
