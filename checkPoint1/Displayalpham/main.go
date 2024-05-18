package main

import "github.com/01-edu/z01"

func main() {
	Displayalpham()
	Displayalrevm()
}

func Displayalpham() {
	for i := 'a'; i <= 'z'; i++ {
		if i%2 == 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}

func Displayalrevm() {
	for i := 25; i >= 0; i-- {
		letter := 'a' + i
		if letter%2 != 0 {
			z01.PrintRune(rune(letter - 32))
		} else {
			z01.PrintRune(rune(letter))
		}
	}
	z01.PrintRune('\n')
}
