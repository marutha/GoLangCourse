package main

import "fmt"

func main() {
	bd := 1982
	years := 0
	for {
		if bd > 2020 {
			break
		}
		years++
		bd++
	}
	fmt.Println(years)
}
