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
	z := map[string]person{
		x.last: x,
		y.last: y,
	}
	fmt.Println(z)

	for k, v := range z {
		fmt.Println(k, v)
		for i, j := range v.favourite {
			fmt.Println(i, j)
		}
	}
}
