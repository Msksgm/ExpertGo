package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)

	ch <- 100

	select {
	case <-ch:
		fmt.Println("received")
	default:
		fmt.Println("block")
		break
	}
}
