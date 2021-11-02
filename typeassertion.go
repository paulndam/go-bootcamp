package main

import (
	"fmt"
	"time"
)

// when  having a value and want to change that value or convert it to another or specific type. We use a type assertion.

//  a type assertion takes in a value and tries to convert it to another specific type.

type Stringer interface{
	String() string
}

type fakeString struct {
	content string
}

// function that implement the Stringer interface.
func (s *fakeString) Stringer() string {
	return s.content
}

func printString(value interface{}){
	switch str := value.(type){
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func timeMap(y interface{}){
	z,ok := y.(map[string]interface{})

	if ok {
		z["updated_at"] = time.Now().Weekday()
	}
}


func main(){
	foo := map[string]interface{}{
		"Henry":56,
	}
	timeMap(foo)
	fmt.Println(foo)

	s := &fakeString{"this is not a string"}
	printString(s)
	printString("Hello all Gophers")
}