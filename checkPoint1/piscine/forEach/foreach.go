package main

import (
	"fmt"
)
// Write a function ForEach that, for an int slice, applies a function on each element of that slice.

func ForEach(f func(int), a []int) {
	for _, v := range(a) {
		f(v)
	}
}


func printElement(x int) {
    fmt.Println("Element:", x)
}

// Function to square the element
func squareElement(x int) {
    fmt.Println("Squared Element:", x*x)
}

// Function to check if the element is even
func checkEven(x int) {
    if x%2 == 0 {
        fmt.Println("Element is even:", x)
    } else {
        fmt.Println("Element is odd:", x)
    }
}

// Function to calculate the factorial of the element
func calculateFactorial(x int) {
    factorial := 1
    for i := 1; i <= x; i++ {
        factorial *= i
    }
    fmt.Println("Factorial of", x, "is", factorial)
}

func main() {
    // Example usage of ForEach with different functions
    numbers := []int{1, 2, 3, 4, 5}

    fmt.Println("Original Slice:", numbers)

    fmt.Println("\nPrinting elements:")
    ForEach(printElement, numbers)

    fmt.Println("\nSquaring elements:")
    ForEach(squareElement, numbers)

    fmt.Println("\nChecking even or odd:")
    ForEach(checkEven, numbers)

    fmt.Println("\nCalculating factorial:")
    ForEach(calculateFactorial, numbers)
}
