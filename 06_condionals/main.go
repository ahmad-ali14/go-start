package main

import (
	"fmt"
)

func main() {

	//if else

	x, y := 5, 10

	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else if y > x {
		fmt.Printf("%d is greater than %d", y, x)
	} else {
		fmt.Printf("%d is equall to  %d", x, y)
	}

	fmt.Println("\n")

	// switch
	color := "gray"

	switch color {
	case "red":
		fmt.Println("red", color)
	case "blue":
		fmt.Println(" not red", color)
	default:
		fmt.Println("not red or blue", color)
	}

}
