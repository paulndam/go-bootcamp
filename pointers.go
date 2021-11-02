package main

import "fmt"

// struct fields can be access thru a struct pointer.

/*

To get pointer of value, use the & sign, infront of the value.
To dereference the pointer, use the * symbol/sign


*/

func main(){

	age := 25

	myAge := &age

	*myAge = 30

	

	fmt.Println(*myAge)
}