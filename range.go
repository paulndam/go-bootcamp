package main

import "fmt"

// range form of a loop iterates over a slice or a map

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

var names = []string{
	"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis",
}

func main() {

	for i, v := range pow {
		if v%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("false")
		}

		fmt.Printf("2 ** %d = %d\n", i, v)
	}

	var maxLen int

	for _, vl := range names {
		if l := len(vl); l > maxLen {
			maxLen = l
		}
	}

	output := make([][]string, maxLen)

	for _, vl := range names {
		output[len(vl)-1] = append(output[len(vl)-1], vl)
	}

	fmt.Printf("%v", output)

}
