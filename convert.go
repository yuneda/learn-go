package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var niali64 int64 = int64(nilai32)
	var nilai8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(niali64)
	fmt.Println(nilai8)

	var name = "Yuneda"
	var e = name[0]
	var eString = string(e)
	fmt.Println(name)
	fmt.Println(eString)
}