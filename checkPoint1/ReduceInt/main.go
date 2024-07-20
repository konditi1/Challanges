package main

import (
	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	PrintNum(result)
	z01.PrintRune('\n')
}

func PrintNum(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// var PrintNumInRune func(n int)

	// PrintNumInRune = func(n int) {
	// 	if n == 0 {
	// 				return
	// 			}
	// 			PrintNumInRune(n / 10)
	// 			z01.PrintRune(rune('0' + (n % 10)))
	// 		}

	PrintNumInRune(n)
}

func PrintNumInRune(n int) {
	if n == 0 {
		return
	}
	PrintNumInRune(n / 10)
	z01.PrintRune(rune('0' + (n % 10)))
}
