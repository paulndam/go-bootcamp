package main

import "fmt"

type Animal struct{
	Dog string
	Lion string
}


func main(){

	// Initialization.

	x := new(Animal)
	y := &Animal{}

	fmt.Println(*x == *y)
}