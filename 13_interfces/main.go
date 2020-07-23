package main

import (
	"fmt"
	"math"
)

// interface
type Shape interface {
	area() float64
}

// struct 1
type Circle struct {
	x, y, r float64
}

//struct 2
type Rectangle struct {
	w, h float64
}

// func of struct1
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// func of struct2
func (r Rectangle) area() float64 {
	return r.w * r.h
}

// func of interface
func getArea(s Shape) float64 {
	return s.area()
}

// main
func main() {
	cr := Circle{1, 2, 3}
	rct := Rectangle{4, 5}

	fmt.Printf("Cr area is %f \n", getArea(cr))
	fmt.Printf("Rct area is %f \n", getArea(rct))

}
