package main

import "fmt"

func main() {
	// slice
	ids := []int{12, 14, 44, 32, 99}

	// loop with range over arrays
	for i, id := range ids {
		fmt.Println(i, id)

	}

}
