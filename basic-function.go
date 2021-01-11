package main

import(
	"fmt"
)

func basicFunc(x int, y int) int{
	return x + y
}

func sameSingleTypeFunc(x, y int) int{
	return x+y
}

func returnTwoVariable(x, y string) (string, string){
	return x, y
}

func main(){
	fmt.Println("2 + 4 is ", basicFunc(2,4))
	fmt.Println("5 + 10 is ", sameSingleTypeFunc(5,10))
	fmt.Println(returnTwoVariable("Hello", "World"))

	a, b := returnTwoVariable("First", "Second")
	fmt.Println(a, b)
}