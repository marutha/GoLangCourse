package main

import "fmt"

func main() {
	s :=
		struct {
			first   string
			friends map[string]int
			fav     []string
		}{
			first: "M",
			friends: map[string]int{
				"murali":    123,
				"mokkarasu": 345,
				"loose":     798,
			},
			fav: []string{"vodka", "beer"},
		}
	fmt.Println(s)
	for k, v := range s.friends {
		fmt.Println(k, v)
	}
}
