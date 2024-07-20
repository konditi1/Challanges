package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	num := Atoi(os.Args[1])
	for i := 1; i < 10; i++ {
		PrintNum(i)
		Print(" x ")
		PrintNum(num)
		Print(" = ")
		PrintNum(i*num)
		z01.PrintRune('\n')
	}
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

func Atoi(s string)int {
	var n int
	for _, num := range s {
		digit := int(num - '0')
		n = 10*n + digit
	}
	return n
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
