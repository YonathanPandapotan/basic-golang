package main

import(
	"fmt"
)

func main(){

	// Variable declaration works almost like kotlin
	// var variable-name variable-type
	var i int
	i = 19
	// var logic bool = false
	var text = "Hello World"

	// We can use walrus operator to create and give variable value without declaring

	walrus := 0.444

	fmt.Println(i)
	fmt.Println(text)
	fmt.Println(walrus)

	// convert or casting
	// Maybe cannot cast numeric to string
	angka := 10
	uintAngka := uint(10)
	floateAngka := float64(textAngka)

	fmt.Println(angka)
	fmt.Println(uintAngka)
	fmt.Println(floateAngka)
}