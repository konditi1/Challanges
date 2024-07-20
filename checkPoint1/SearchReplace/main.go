package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("args must be 3")
		return
	}
	word := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]
	str := ""
	for _, c := range word {
		if string(c) == search {
			str += replace
		} else {
			str += string(c)
		}
	}

	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
