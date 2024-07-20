package main

import (
	"github.com/01-edu/z01"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
	
	z01.PrintRune('\n')
}

func Chunk(slice []int, size int) {
	if size == 0 {
		z01.PrintRune('\n')
		return
	}
	arr := [][]int{}

	for i := 0; i < len(slice); i += size {
		end := size + i
		if end > len(slice) {
			end = len(slice)
		}
		arr = append(arr, slice[i:end])

	}
	PrintArr(arr)
}

func PrintArr(arr [][]int) {
	z01.PrintRune('[')
	for i, outer := range arr {
		z01.PrintRune('[')
		for j, inner := range outer {
			PrintNum(inner)
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

func PrintNum(n int) {
	if n == 0 {
		z01.PrintRune('0')
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
		PrintNumInRune(n / 10)
		z01.PrintRune(rune('0' + (n % 10)))
	}
	PrintNumInRune(n)
}
