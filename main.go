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
		fmt.Println(err)
	}

	// user := User{
	// 	Username: "root",
	// 	Password: "toor",
	// }
	// err = conduit.From("users").Insert(user)
	// if err != nil {
	// 	panic(err)
	// }

	results, err := conduit.From("users").Where(func(entry Entry) bool {
		return entry["username"] == "root"
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	view := results.Select("username")
	for i, row := range view.List() {
		fmt.Println(i, row)
	}
}
