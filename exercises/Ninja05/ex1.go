package main

import "fmt"

type person struct {
	first     string
	last      string
	favourite []string
}

func main() {
	x := person{
		first:     "first1",
		last:      "last1",
		favourite: []string{"vanilla", "chocolate"},
	}
	y := person{
		first:     "first2",
		last:      "last2",
		favourite: []string{"vanilla2", "chocolate"},
	}
	fmt.Println(x)
	fmt.Println(y)
}
