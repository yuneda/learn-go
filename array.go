package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Achmad"
	names[1] = "Yuneda"
	names[2] = "Alfajr"
	// names[3] = "Alfat"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		1,
		4,
		8,
	}

	fmt.Println(values)
	fmt.Println(len(values))

	var SepuluhData [10]string

	fmt.Println(SepuluhData)
	fmt.Println(len(SepuluhData))
}
