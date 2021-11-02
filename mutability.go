package main

import "fmt"

// in Go only constants are immutable.

type Artist struct{
	Name,Genre string
	songs int
}

func newsongs(a *Artist) int {

	a.songs++
	return a.songs
}

func main(){
	cb := &Artist{Name: "chris brown",Genre: "R&B",songs: 190}
	fmt.Printf("%s have released their %dth song\n",cb.Name,newsongs(cb))
	fmt.Printf("%s has a total of %d songs", cb.Name,cb.songs)
}