package main

import "testing"

func Greet(name string) string {
	greet := "Hello, " + name + "!"
	return greet
}

func TestGreet(t *testing.T) {
	result := Greet("John Doe")
	expected := "Hello, John Doe!"

	if result != expected {
		t.Errorf("Unexpected greeting: %s", result)
	}
}
