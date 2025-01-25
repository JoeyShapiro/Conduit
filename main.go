package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	conduit, err := New("test.db")
	if err != nil {
		fmt.Println(err)
	}

	conduit.From("benchmarks").Execute()
}
