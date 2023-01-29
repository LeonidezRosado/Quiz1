package main

import (
	"fmt"
	"math"
)


//1)creating a type called "triangle" with two fields named base and height (type = float64)
type triangle struct{
	base float64
	height float64
}

//2)creating a method on type triangle named area that calculates and returns area 
func (t triangle) Area() float64{
	return 1/2 * (t.base * t.height)
}

func main() {

}