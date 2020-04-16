package main

import "fmt"

func main() {
	fmt.Println("hello this is the code")
	foo()
	fmt.Println("This is the second line")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf(" %d \n", i)
		}
	}
}

func foo() {
	fmt.Println("I am in foo")
}
