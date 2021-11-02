package main

import "fmt"

// constants are declared like variables.
// constants can only be characters, strings, bool and numeric values.
// they can't be defined using the := operator

const pi = 3.14

const (
	StatusOk = 200
	StatusCreated = 201
	StatusNonAuthoritativeInfo = 203
	StatusNoContent = 204
	StatusResetContent = 205
	StatusPartialContent = 206
)

const (
	Pi = 3.14
	Truth = false
	Big = 1 << 62
	Small = Big >> 61
)

func main(){
	const Greeting =""
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
	fmt.Println(Small)
	
}
