package main

import "fmt"

func main() {
	i := interface{}("Go Expert")

	s := i.(string)
	fmt.Println(s)

	// n := i.([]byte)
	// fmt.Println(n)

	n, ok := i.([]byte)
	fmt.Println(n, ok)
}
