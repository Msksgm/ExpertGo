package main

import "fmt"

type Crier interface {
	Cry() string
}

type Duck struct{}

func (d Duck) Cry() string {
	return "Quack"
}

var c Crier = Duck{}

func main() {
	fmt.Println(c.Cry())
}
