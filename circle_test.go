package main

import (
	"testing"
)

func TestCircle_area(t *testing.T) {
	got := circle {
		radius: 3,	
	}

	got1:= got.area()
	expected := 28.274333882308138

	if got1 != expected {
		t.Errorf("got %v expected %v", got1, expected)
	}
}

func TestCircle_perimeter(t *testing.T) {
	got := circle {
		radius: 3,
	}

	got1 := got.perimeter()
	expected := 18.84955592153876

	if got1 != expected {
		t.Errorf("got %v expected %v", got, expected)
	}
}