package main

import(
	"fmt"
)

func main(){
	// Basic for
	var sum int = 0
	for i:=0; i<10 ; i++{
		sum += i
	}
	fmt.Println(sum)

	// While (for is go's while)
	num2 := 0
	for num2 <5{

		num2+=1
	}
	fmt.Println(num2)
}