package main

import (
	"math"
)

//6)creating a type named circle that has one field named 'radius' of type float64
type circle struct{
	radius float64
}

//7)creating a method on type circle named "area" that calculates and returns area 
func (c circle) area() float64 {
	var y float64 = 2
	return  math.Pi * (math.Pow(c.radius,y))
}

//8 creating a method on type circle named "perimeter" that calculates and returns perimeter
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

