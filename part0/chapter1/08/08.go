package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}

	fmt.Println(src, len(src), cap(src))
	dst := make([]int, len(src))
	copy(dst, src)

	fmt.Println(dst, len(dst), cap(dst))
}
