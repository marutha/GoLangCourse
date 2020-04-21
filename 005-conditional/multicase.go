package main

import "fmt"

func main() {
	x := "Bond"
	switch x {
	case "Maru", "Bond", "Match":
		fmt.Println("This should print")
	case "m":
		fmt.Println("this should not print")
	default:
		fmt.Println("This should print")
	}
}
