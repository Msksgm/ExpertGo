package main

import "fmt"

type (
	ErrNoSuchEntity   struct{ error }
	ErrConflictEntity struct{ error }
)

func main() {
	do := func() error {
		return &ErrConflictEntity{}
	}

	switch do().(type) {
	case *ErrNoSuchEntity:
		fmt.Println("error no such entity")
	case *ErrConflictEntity:
		fmt.Println("error confilict entity")
	default:
		fmt.Print("unknown error")
	}
}
