package main

import (
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	fmt.Println("hello world")

	conduit, err := New("test.db")
	if err != nil {
		fmt.Println(err)
	}

	err = conduit.Create("users")
	if err != nil {
		panic(err)
	}

	user := User{
		Username: "root",
		Password: "toor",
	}
	err = conduit.From("users").Insert(user)
	if err != nil {
		panic(err)
	}

	results, err := conduit.From("users").Where(func(entry any) bool { return true }).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(results)
}
