package main

import (
	"fmt"
	"time"
)

func main() {
	doneCh := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		doneCh <- 100
	}()

	for {
		select {
		case <-time.Tick(1 * time.Second):
			fmt.Println("waiting...")
		case <-doneCh:
			fmt.Println("Done")
			return
		}
	}
}
