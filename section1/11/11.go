package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}

	for i := len(src)/2 - 1; i >= 0; i-- {
		opp := len(src) - 1 - i
		src[i], src[opp] = src[opp], src[i]
	}
	fmt.Println(src)

	for left, right := 0, len(src)-1; left < right; left, right = left+1, right-1 {
		src[left], src[right] = src[right], src[left]
	}
	fmt.Println(src)
}
