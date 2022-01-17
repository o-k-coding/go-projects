package main

import "fmt"



func main() {
	// Map with keys of string and values of string
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)

	// have to use make to get type inferance
	ip := make(map[string]string)
	fmt.Println(ip)

	colors["blue"] = "#4b7fff"

	delete(colors, "red")
	fmt.Println(colors)

	// readIPsFromFile("./ip-address-10k.txt")
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("hex code for ", color, "is", hex)
	}
}

// maps vs structs
// maps are reference type, not value type like struct
// maps are easy to loop over, structs are not
// structs can have any value types, maps must have all keys and values the same type

// Maps should be used to represent a collection of closely related properties

