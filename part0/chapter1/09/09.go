package main

import "fmt"

func main() {
	src1, src2 := []int{1, 2}, []int{1, 2, 3, 4, 5}

	dst := append(src1, src2...)
	fmt.Println(dst)
}
