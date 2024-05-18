package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	args := os.Args[1:]
	
	for _, v := range args[0] {
		z01.PrintRune(rune(v))
	}
	z01.PrintRune('\n')
}