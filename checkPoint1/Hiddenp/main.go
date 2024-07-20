package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	z01.PrintRune(Hiddenp(s1, s2))
	z01.PrintRune('\n')
}

func Hiddenp(s1, s2 string) rune {
	var str string
	for _, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				str += string(c1)
				s2 = s2[j+1:]
				break
			}
		}
	}
	if str == s1 || s1 == "" {
		return '1'
	}
	return '0'
}
