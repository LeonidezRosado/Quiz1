package main

import (
	"testing"
)
//10)write a test for square function 
func TestSquare(t *testing.T) {
	got1,got2 := square(10)
	expected1, expected2 := 100,40
	if got1  != float64(expected1) && got2 != float64(expected2) {
		t.Errorf("got %v expected %v, \n got %v expected %v", got1,expected1,got2,expected2)
	}
}
