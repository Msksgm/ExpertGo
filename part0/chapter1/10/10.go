package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}

	i := 2
	dst := append(src[:i], src[i+1:]...)
	fmt.Println(dst)

	src = []int{1, 2, 3, 4, 5}
	dst = src[:i+copy(src[i:], src[i+1:])]
	fmt.Println(dst)
}
