package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpYuneda NoKTP = "1234567890"
	marriedStatus := Married(true)
	fmt.Println(noKtpYuneda)
	fmt.Println(marriedStatus)
}
