package main

import "fmt"

func main() {
	x := map[string]string{
		"hello": "world",
		"maru":  "s",
	}

	fmt.Println(x)
	for i, v := range x {
		fmt.Printf("key:%v value: %v\n", i, v)
	}
}
