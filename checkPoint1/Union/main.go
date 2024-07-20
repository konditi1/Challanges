package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args1 := os.Args[1]
	args2 := os.Args[2]
	str := ""

	for _, c := range args1 {
		for _, r := range args2 {
			if r == c && !IsPresent(str, c) {
				str += string(c)
			}
		}
	}
	Print(str)
}

func IsPresent(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
