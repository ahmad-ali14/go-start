package main

import "fmt"

func main() {

	// define a map
	emails := make(map[string]string)

	// assign
	emails["me"] = "me@maile.me"

	//print
	fmt.Println(emails)

	// delete from map
	delete(emails, "me")

	// shorthand,
	jobs := map[string]string{"ahmad": "devedloper", "before": "pharmacist"}

	fmt.Println(jobs)
}
