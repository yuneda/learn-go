package main

import "fmt"

func main() {
	var name string
	name = "John Wick"
	fmt.Println(name)

	name = "John Wick 2"
	fmt.Println(name)

	var friendName = "Joko"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "John"
		lastName  = "Wick"
	)
	fmt.Println(firstName[0])
	fmt.Println(lastName)
}
