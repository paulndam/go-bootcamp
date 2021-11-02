package main

import (
	"fmt"
	"time"
)

// struct is a collection of fields or properties.
/*
A struct can be considered as a class to supports composition but not inheritance.

*/

type Car struct {
	Make string
	Model string
	Year int 
	Trim string 
	Price float64 
	Engine string
	Release time.Time

}

func main(){
	event := Car{
		Make: "Mercedes Benz",
		Model: "S Class",
		Year: 2021,
		Trim: "AMG 63 S + 4MATIC",
		Price: 113565.78,
		Engine: "V8 Twin turbo +",
	}

	fmt.Println(event)

	// accessing fields using the dot notation.

	event.Release = time.Now()

	fmt.Printf("The release day is on %s, location will be at Boston",event.Release)
}