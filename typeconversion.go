package main

import "fmt"

// the expression T(v) converts the value to the type T.

var i = 42
var f float64 = float64(i)
var u uint = uint(f)



// Or simple a type assertion.

func main(){
	a := 40
	b := float64(a)
	c := uint(b)

	fmt.Println(a,b,c)

	fmt.Println(i,f,u)

}