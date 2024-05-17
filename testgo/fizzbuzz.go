package main

import (
	"fmt"
)

func FizzBuzz(n int) string {
	if (n % 3 == 0){
		return Fizz(n)
	}
	return ToString(n)
}
func Fizz(n int) string{
	return "Fizz"
}
func ToString(n int) string{
	return fmt.Sprintf("%v", n)
}