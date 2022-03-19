package main

import "fmt"

func main() {
	i := interface{}("Go Expert")

	switch i.(type) {
	case int, int8, int16, int32, int64:
		fmt.Println("This is integer,", i)
	case string:
		fmt.Println("This is string,", i)
	default:
		fmt.Printf("This is uknown type, %T\n", i)
	}
}
