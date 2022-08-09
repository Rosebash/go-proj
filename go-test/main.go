package main

import "fmt"

type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func main() {
	a := Book{
		Name:        "The Last Lion",
		Author:      "JJ Redick",
		Publication: "hee",
	}
	fmt.Println(a)
}
