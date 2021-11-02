package main

import (
	"fmt"
	"strings"
)

type Tips struct {
	Amount, Taxes float64
}

var m map[string]Tips

// word count
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {

	players := map[string]int{
		"messi":   6,
		"ronaldo": 5,
		"iniesta": 1,
		"neymar":  2,
	}

	fmt.Printf("%v", players)

	m = make(map[string]Tips)
	m["Etoo"] = Tips{328.89, 54.90}

	fmt.Println(m["Etoo"])

	// mutating maps.

	// insert element into map.
	// m[key] = element
	// retrieve.
	// element = m[key]
	// delete
	// delete(m,key)
	// check if key is present
	// element,ok = m[key]

	fmt.Println(WordCount("hello guys hello"))

}
