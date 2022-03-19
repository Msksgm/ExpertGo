package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// メソッドを値として使う
	fv := Hex(1024).String
	fmt.Println(fv())

	// メソッドを式として使う
	fe := Hex.String
	fmt.Println(fe(1024))
}
