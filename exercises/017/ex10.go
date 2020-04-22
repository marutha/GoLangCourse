package main

import "fmt"

func main() {
	x := map[string]string{
		"hello": "world",
		"maru":  "s",
	}
	x["kavitha"] = "subramoney"
	fmt.Println(x)
	for i, v := range x {
		fmt.Printf("key:%v value: %v\n", i, v)
	}
	delete(x, "maru")
	fmt.Println()
	for i, v := range x {
		fmt.Printf("key:%v value: %v\n", i, v)
	}
}
