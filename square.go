package main

import (
	"math"
)

//9)creating a go function named square that accepted 
func Square(side float64) (float64, float64) {
	var y float64 = 2
	return math.Pow(side,y), side + side + side + side
}