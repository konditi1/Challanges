package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	num := Atoi(os.Args[1])
	num2 := Atoi(os.Args[2])
	gcd := Gcd(num, num2)
	Print(gcd)
	z01.PrintRune('\n')
}

func Gcd(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Atoi(s string) int {
	var n int
	for _, num := range s {
		dig := int(num - '0')
		n = n*10 + dig
	}
	return n
}

func Print(i int) {
	if i == 0 {
		z01.PrintRune('0')
		return
	}

	// Helper function to print digits recursively
	var printDigits func(int)
	printDigits = func(n int) {
		if n == 0 {
			return
		}
		printDigits(n / 10)
		z01.PrintRune(rune('0' + (n % 10)))
	}

	printDigits(i)
}
