package main

import (
	"fmt"
	"strings"
)

var boolMap = map[bool]int{false: 0, true: 1} 
func FizzBuzz(n int) string {
	s := Fizz(n)
	if(len(s) == 0){
		s = ToString(n)
	}
	return s
}
func Fizz(n int) string{
	return strings.Repeat("Fizz", boolMap[n%3==0])
}
func Buzz(n int) string{
	return "Buzz"
}
func ToString(n int) string{
	return fmt.Sprintf("%v", n)
}