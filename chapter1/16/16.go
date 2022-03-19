package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m == nil)

	fmt.Println(len(m))

	v, ok := m["John"]
	fmt.Println(v, ok)

	delete(m, "Richard")
	m["John"] = 42
}
