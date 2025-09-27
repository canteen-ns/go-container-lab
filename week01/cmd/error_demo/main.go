package main

import (
	"errors"
	"fmt"
	myerrors "github.com/canteen_ns/go-container-lab/week01/pkg/errors"
	"github.com/canteen_ns/go-container-lab/week01/pkg/iface"
)

func main() {
	a := 42
	iface.ShowEface(a)

	err := myerrors.Top()
	if err != nil {
		fmt.Println("Error chain:")
		fmt.Println(err)

		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("unwrap -> %v\n", e)
		}
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	panic("demo panic")
}
