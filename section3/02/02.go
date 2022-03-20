package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case ch <- 100:
		fmt.Println("sent")
	default:
		fmt.Println("block")
	}
}
