package main

import "fmt"


func main(){


	a := []int{1,5,6,8,9,13}
	fmt.Println(a)

	// slicing a slice.
	fmt.Println(a[1:4])

	// missing low index implies index 0.
	fmt.Println(a[:3])

	// missing high index implies len().
	fmt.Println(a[4:])


	// making slice using make()

	cities := make([]string,2)
	cities[0] = "baltimore"
	cities[1] = "new york"

	fmt.Printf("%q",cities)

	// appending to a slice.
	shoe := []string{}
	shoe = append(shoe,"nike","adidas","puma")
	fmt.Printf("%q",shoe)

	// append a slice to another slice using ellipsis.
	car := []string{"toyota","gmc"}
	cars := []string{"mb","audi"}
	car = append(car,cars...)
	fmt.Printf("%q",car)

	// check lenght of slice.
	fmt.Println(len(car))



}