package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	// Chunk([]int{}, 10)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 99}, 0)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 99}, 3)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	// Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
	// PrintNumInRune(-17)
	fmt.Println(Repeat("abd"))
}

func Chunk(slice []int, size int) {
	var arr [][]int
	if size == 0 {
		z01.PrintRune('\n')
		return
	}
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		arr = append(arr, slice[i:end])
	}
	z01.PrintRune('[')
	for i, outer := range arr {
		z01.PrintRune('[')
		for j, inner := range outer {
			PrintNumInRune(inner)
			if j != len(outer)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')
		if i != len(arr)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

func PrintNumInRune(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Helper function to print digits recursively
	var printDigits func(int)
	printDigits = func(n int) {
		if n == 0 {
			return
		}
		printDigits(n / 10)
		z01.PrintRune(rune('0' + (n % 10)))
	}

	printDigits(n)
}

func Repeat(s string) (new string) {
	length := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			length = int(c - 'a' + 1)
		} else if c >= 'A' && c <= 'Z' {
			length = int(c - 'A' + 1)
		}
		for length > 0 {
			new += string(c)
			length--
		}
	}
	return
}
