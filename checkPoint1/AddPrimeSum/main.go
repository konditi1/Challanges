package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	num, err := Atoi(os.Args[1])
	if err != "nil" {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	PrintNum(AddPrimeSum(num))
	z01.PrintRune('\n')

}

func AddPrimeSum(n int) int {
	var sum int
	for i := n; i > 1; i-- {
		if IsPrime(i) {
			sum += i
		}
	}
	return sum

}

func Atoi(s string) (int, string) {
	var n int
	for _, c := range s {
		if c < '0' || c > '9' || c == '-' {
			return 0, "ERROR: converting non int"
		}
		digit := int(c - '0')
		n = n*10 + digit
	}
	return n, "nil"
}
func PrintNum(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var PrintNumInRune func(n int)
	PrintNumInRune = func(n int) {
		if n == 0 {
			return
		}
		PrintNumInRune(n / 10)
		z01.PrintRune(rune('0' + (n % 10)))
	}
	PrintNumInRune(n)
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
