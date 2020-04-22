package main

import "fmt"

func main() {
	bd := 1982
	years := 0
	for bd < 2020 {
		years++
		bd++
	}
	fmt.Println(years)
}
