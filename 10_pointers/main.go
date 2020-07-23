package main

import "fmt"

func main() {

	// poiter
	a := 4
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T", b)

	// use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// change val with pointers
	*b = 9
	fmt.Println(a, b)

}
