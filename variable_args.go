package main

// Importing required packages
import (
	"fmt"
)

// Variadic function to return
// the sum of the numbers
func add(num ...int) int {
	sum := 0
	for j := range num {
		sum += j
	}
	return sum
}

func main() {
	fmt.Println("Sum =", add())
}
