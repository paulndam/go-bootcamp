package main

import "fmt"

// declairing one type parameter for both params.
func add(x,y int) int {
	return x + y
}

func location(city string)(country,continent string){
	switch city {
	case "Baltimore", "MD","North America":
		country, continent = "U-S-A", "North America" 
	case "Toky","Asia":
		country, continent = "Japan","East Asia"
	default:
		country,continent = "unknown", "unknown"
	}
	return 
}

func location2(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
			region, continent = "California", "North America"
	case "New York", "NYC":
			region, continent = "New York", "North America"
	default:
			region, continent = "Unknown", "Unknown"

} 
return region, continent
	}

func main(){
	fmt.Println(add(50,87))

	region, continent := location2("Santa Monica")
	fmt.Printf("Matt lives in %s, %s", region, continent)
	
	country, continent := location("East Asia")
	fmt.Printf("Nagasaki lives in %s, %s", country,continent)
}