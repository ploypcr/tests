package main

import (
	"fmt"
)

var boolMap = map[bool]int{false: 0, true:1}

func FizzBuzz(n int) string {
	return ToString(n)
}

func ToString(n int) string{
	return fmt.Sprintf("%v", n)
}