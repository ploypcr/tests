package main

import (
	"fmt"
)

func FizzBuzz(n int) string {
	s := ToString(n)
	return s
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