package main

import (
	"fmt"
)

func FizzBuzz(n int) string {
	if (n % 3 == 0){
		return Fizz(n)
	}
	if (n % 5 == 0){
		return Buzz(n)
	}
	return ToString(n)
}
func Fizz(n int) string{
	return "Fizz"
}
func Buzz(n int) string{
	return "Buzz"
}
func ToString(n int) string{
	return fmt.Sprintf("%v", n)
}