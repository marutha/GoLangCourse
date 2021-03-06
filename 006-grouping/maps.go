package main

import "fmt"

func main() {
	m := map[string]int{
		"james":  1,
		"killer": 2,
	}

	fmt.Println(m["james"])
	fmt.Println(m["fun"])

	if v, ok := m["killer"]; ok {
		fmt.Println("here is the value", v)
	}

	m["hello"] = 455
	fmt.Println()
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println()
	delete(m, "killer")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
