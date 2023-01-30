package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	got := triangle{
		base: 10,
		height: 12,
	}
	got1 := got.Area()
	expected := 60
	

	if got1 != float64(expected) {
		t.Errorf("got %v expected %v", got, expected)
	}

}



func TestPerimeter(t *testing.T) {
	got := triangle{
		base: 10,
		height: 12, 
	}
	got1 := got.perimeter()
	expected1 := 37.62049935181331

	if got1 != expected1 {
		t.Errorf("got %v expected %v", got, expected1)
	}

}