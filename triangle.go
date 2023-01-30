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
	return .5 * (t.base * t.height)
}

//3)creating a method on type triangle named perimeter that calculates and returns perimeter
func (t triangle) perimeter() float64 {
	var y float64 = 2
	hypo := math.Sqrt(math.Pow(t.base,y) + math.Pow(t.height,y))

	return hypo + t.base + t.height
}

func main() {
	//4)creating a variable of type triangle
	Trian1 := triangle{
		base: 10,
		height: 12,
	}

	//5)calling the area and permeter methods on that triangle
	fmt.Println(Trian1.Area())
	fmt.Println(Trian1.perimeter())

	circ1 := circle {
		radius: 3,
	}

	fmt.Println(circ1.area())
	fmt.Println(circ1.perimeter())

	//calling the square function
	area, perimeter := square(10)
	fmt.Println("Area of a square:",area)
	fmt.Println("perimeter of a square:",perimeter)
}