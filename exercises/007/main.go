package main

import "fmt"

func main() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 45)
	d := (45 != 45)
	e := (45 > 45)
	f := (45 < 45)
	fmt.Println(a, b, c, d, e, f)
}
