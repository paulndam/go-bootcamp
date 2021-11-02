package main

import "fmt"

// declairing a list of variables.

var (
	name string
	age int 
	location string
)

// second option.
var (
	name, location string
	age int
)

// variables can also be declared one by one.
var name string
var age int
var location string 

// a var delcaration can include initializers.
var (
	name = "same"
	age = 34
	location = "Portugal"

)

// variables can also be initialized on the same line.
var (name,location,age = "berry red","Malibu",43)

// in a function the short assignment statement can be used in the place of a var with implicit type.
func main(){
	name, location := "terry miles","Ohio"
	age := 23

	fmt.Println("%s (%d) of %s", name,age,location)

	action := func(){
		// do something.
		fmt.Println("i am a function but a variable have been assigned to hold me")
	}

	action()
}

