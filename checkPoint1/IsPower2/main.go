package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num := Atoi(os.Args[1])
	Print(Ispower2(num))

}

func Atoi(s string) int {
	var n int
	for _, num := range s {
		dig := int(num - '0')
		n = n*10 + dig
	}
	return n
}

func Ispower2(n int) string {
	for n > 1 {
		if n % 2 != 0 {
			return "false"
		}
		n /= 2
	}
	return "true"
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}