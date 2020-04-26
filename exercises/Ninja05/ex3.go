package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	car := sedan{
		vehicle: vehicle{
			door:  4,
			color: "red",
		},
		luxury: true,
	}
	mack := truck{
		vehicle: vehicle{
			door:  2,
			color: "blue",
		},
		fourWheel: true,
	}
	fmt.Println(car)
	fmt.Println(mack)
}
