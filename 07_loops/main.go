package main

import "fmt"

func main() {

	// for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	//for loop shorthand
	for x := 0; x < 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("FuzzBuzz")
			continue
		}
		if x%3 == 0 {
			fmt.Println("Buzz")
			continue
		}
		if x%5 == 0 {
			fmt.Println("Fuzz")
			continue
		}

		fmt.Println(x)
	}
}
