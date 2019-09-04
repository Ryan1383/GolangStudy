package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {

	author := Author{Name: "Ellot Forbes", Age: 25, Developer: true}
	book := Book{Title: "Leaning Concurrency in Python", Author: author}

	// fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
