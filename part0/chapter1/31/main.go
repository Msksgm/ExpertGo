package main

import "fmt"

func main() {
	msg := "Go Expert"

	bs := []byte(msg)
	fmt.Println(bs)

	s := string(bs)
	fmt.Println(s)
}
