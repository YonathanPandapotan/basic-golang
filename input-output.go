package main

import(
	"fmt"
)

func add(x int, y int) int{
	return x+y
}

func main(){
	fmt.Println("2 + 4 is ", add(2,4))
}