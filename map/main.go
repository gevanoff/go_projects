package main

import "fmt"

func main() {
	colors := map[string]string{ // map where keys are type string and values are type string
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	/* var colors map[string]string // Alternate way to declare a map
	colors := make(map[string]string) // Equivalent to above, using built in "make" function
	colors["white"] = "#ffffff"       // Add a value to a map at key "white"
	delete(colors, "white")           // Delete key/value pair identified by "white"
	fmt.Println(colors)
	*/

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c { // color and hex correspond to key and value here
		fmt.Println("Hex code for", color, "is", hex)
	}
}

/*
With a Map, all keys must be of the same type and all values must be of the same type.
Structs can have a variety of different types for property values.

A map is a reference type while a struct is a value type (re: need for use of pointers)

A struct needs to have all properties defined ahead of time, while a map can add and remove key/value pairs dynamically

A struct doesn't support indexing so can't iterate over them.
*/
