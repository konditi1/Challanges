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
	Print(PigLatin(word))
}

func PigLatin(s string) string {

	str := ""
	isvowel := map[rune]struct{}{
		'a': {}, 'A': {},
		'e': {}, 'E': {},
		'i': {}, 'I': {},
		'o': {}, 'O': {},
		'u': {}, 'U': {},
	}

	for i, c := range s {
		if _, ok := isvowel[c]; ok {
			if i == 0 {
				str = s[i:] + "ay"
			} else  {
				str = s[i:] + s[:i] + "ay"
			} 
		} else {
			str = "No vowels"
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
