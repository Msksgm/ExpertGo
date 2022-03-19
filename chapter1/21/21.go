package main

import "fmt"

type Chip struct {
	Number int
}

type Card struct {
	string
	Chip
	Number int
}

func (c *Chip) Scan() {
	fmt.Print(c.Number)
}

func main() {
	c := Card{
		string: "Credit",
		Chip: Chip{
			Number: 4242424242,
		},

		Number: 5454545454,
	}

	// Card には　Scan メソッドがないため、 c.Chip.Scan()が実行される
	c.Scan()
	// Scan メソッドのレシーバは、Card ではなく Chip である
}
