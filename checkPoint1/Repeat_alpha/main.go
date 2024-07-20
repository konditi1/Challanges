package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	word := os.Args[1]
	Print(RepeatAlpha(word))
}

func RepeatAlpha(s string) string {
	str := ""
	var index int
	for _, c := range s {
		if  c >= 'a' && c <= 'z' {
			index = int(c-'a'+1)
		} else if  c >= 'A' && c <= 'Z' {
			index = int(c-'A'+1)
		} else {
			str += string(c)
		}

		for index > 0 {
			str += string(c)
			index--
		}
	}
	return str
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}