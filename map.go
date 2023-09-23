package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Achmad Yuneda Alfajr",
		"Address": "Jl. Pahlawan No.5",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["Address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Achmad Yuneda Alfajr"
	book["ups"] = "Salah"

	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(len(book))

	fmt.Println(book)
}
