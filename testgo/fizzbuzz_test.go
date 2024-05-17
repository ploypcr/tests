package main

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1

	got := FizzBuzz(input)

	want := "1"
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzWhenInputCanModBy3(t * testing.T){
	input := 3

	got := FizzBuzz(input)

	want := "Fizz"
	if got != want {
		t.Errorf("got % q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnBuzzWhenInputCanModBy5(t * testing.T){
	input := 5

	got := FizzBuzz(input)

	want := "Buzz"
	if got != want {
		t.Errorf("got % q but want %q", got, want)
	}
}

func TestFizzBuzzShouldReturnFizzBuzzWhenInputCanModBy3And5(t * testing.T){
	input := 15

	got := FizzBuzz(input)

	want := "FizzBuzz"
	if got != want {
		t.Errorf("got % q but want %q", got, want)
	}
}
