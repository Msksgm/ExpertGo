package main

import "fmt"

type Crier interface {
	Cry() string
}

type ParrotFunc func() string

func (p ParrotFunc) Cry() string {
	return p()
}

var c Crier = ParrotFunc(func() string {
	return "Squawk"
})

func main() {
	fmt.Println(c.Cry())
}
