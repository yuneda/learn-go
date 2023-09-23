package main

import "fmt"

func main() {
	name := "Yuneda"
	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Yuneda":
		fmt.Println("Hello Yuneda")
	default:
		fmt.Println("Kenalan dulu")
	}

	// switch with short statement
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// switch without condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
