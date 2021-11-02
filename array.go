package main

import "fmt"

// declares an array of 10 integers.
var a [10]int 

func main(){
	var b [2] string
	b[0]= "hello"
	b[1]= "mama"

	fmt.Println(b[0],b[1])
	fmt.Println(b)


	// setting array entries.
	a := [2]string{"hi","kenya"}
	fmt.Printf("%q",a)

	// using ellipsis in order to use an implicit length.
	x := [...]string{"hi ","world","let's play soccer"}
	fmt.Printf("%q",x)


	// multi-dimensional array.
	 var y  [2][3] string
	 for i := 0; i < 2; i++{
		 for j := 0; j <  3; j++{
			 y[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		 }

	 }
	 fmt.Printf("%q",y)

}