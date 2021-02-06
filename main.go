package main

import (
	"fmt"
)

func fizzBuzz(input int) string {
	word := fmt.Sprintf("%d", input)

	isZeroWhenMod3 := input%3 == 0
	if isZeroWhenMod3 {
		word = "fizz"
	}

	isZeroWhenMod5 := input%5 == 0
	if isZeroWhenMod5 {
		word = "buzz"
	}

	if isZeroWhenMod3 && isZeroWhenMod5 {
		word = "fizz buzz"
	}

	return word
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
