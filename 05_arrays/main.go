package main

import (
	"fmt"
)

func main() {

	// arrays
	var fruits [2]string

	// assign
	fruits[0] = "Apple"
	fruits[1] = "banana"

	//shorthand,
	langs := [3]string{"js", "py", "go"}

	//slices, you don't need to expect number of elements in an array before hand.
	langSlice := []string{"js", "py", "go"}

	fmt.Println(fruits)
	fmt.Println(langs)

	//lenght
	fmt.Println(len(langSlice))

	//range
	fmt.Println(langs[1:3])

}
