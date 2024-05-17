package main

import (
	"fmt"
	"strings"
)

var boolMap = map[bool]int{false: 0, true: 1} 
func FizzBuzz(n int) string {
	s := Fizz(n)
	s += Buzz(n)
	s += ToString(n)
	return s
}
func Fizz(n int) string{
	return strings.Repeat("Fizz", boolMap[n%3==0])
}
func Buzz(n int) string{
	return strings.Repeat("Buzz", boolMap[n%5==0])
}
func ToString(n int) string{
	var res = fmt.Sprintf("%v", n)
	return strings.Repeat(res, boolMap[n%3 != 0 && n%5 != 0])
}