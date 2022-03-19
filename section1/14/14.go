package main

import "fmt"

func main() {
	mapEmpty := map[string]int{}
	fmt.Println(mapEmpty)
	mapMake := make(map[string]int)
	fmt.Println(mapMake)
	mapCap := make(map[string]int, 10)
	fmt.Println(mapCap)
}
