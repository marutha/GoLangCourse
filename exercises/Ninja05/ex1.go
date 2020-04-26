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
	fmt.Println(x.first)
	fmt.Println(x.last)
	for i, v := range x.favourite {
		fmt.Println(i, v)
	}
	fmt.Println(y)
	for i, v := range y.favourite {
		fmt.Println(i, v)
	}
}
