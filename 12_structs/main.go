package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	name         string
	city, gender string
	age          int
}

// greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, I'm " + p.name + ", and my age is " + strconv.Itoa(p.age)
}

// hasBirthday(pointer reciever)
func (p *Person) hasBirthday() {
	p.age += 10
}

func main() {

	// init Person
	person1 := Person{name: "Ahmad", age: 27, city: "London", gender: "m"}

	// shorthand
	person2 := Person{"Ahmad", "London", "m", 33}

	fmt.Println(person1, person2)
	person1.age++
	fmt.Println(person1.age)
	person1.hasBirthday()
	fmt.Println(person1.greet())
}
