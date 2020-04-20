package main

import "fmt"

var x int = 42
var y string = "Maru"
var z bool = true

func main() {
	s := fmt.Sprintf("%d\t%v\t%v", x, y, z)
	fmt.Println(s)
}
