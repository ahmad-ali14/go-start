package main

import (
	"fmt"
)

func main() {
	// var name type

	// explicit declaration
	var me string = "Ahmad"

	// short hand, only works inside functions
	age := 27

	// shorthand, for multiple vars
	where, when := "London", 2020

	// const vars
	const tired = true

	// print to the terminal
	fmt.Println("I'm ", me, age)
	fmt.Println(when, where)

	// if statments
	if tired {
		fmt.Println(me + " is tired")
	}

	// get the type of a var
	fmt.Printf("%T\n", age)
}
