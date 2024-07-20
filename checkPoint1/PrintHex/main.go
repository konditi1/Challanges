package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	num, err := Atoi(os.Args[1])
	if err != "nil" {
		Print(err)
		return
	}
	Print(IntToHex(num))
}

func IntToHex(n int) string {
	hex := "0123456789abcdef"
	str := ""

	if n == 0 {
		return "0"
	}

	for n > 0 {
		rem := n % 16
		str = string(hex[rem]) + str
		n /= 16
	}
	return str
}

func Atoi(s string) (int, string) {
	var n int
	if s == "0" {
		return 0, "nil"
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, "ERROR"
		}
		digit := int(c - '0')
		n = n*10 + digit
	}
	return n, "nil"
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
