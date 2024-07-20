package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{-1, -2, -3}
	ac := -93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()
	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, table []int, ac int) {
	for _, num := range table {
		ac = f(ac, num)
	}
	PrintNum(ac)
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}

func PrintNum(n int) {
	if n == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
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

		PrintNumInRune(n/10)
		z01.PrintRune(rune('0'+(n%10)))
	}
	PrintNumInRune(n)
	z01.PrintRune('\n')
}
