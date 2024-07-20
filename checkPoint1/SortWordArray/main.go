package main

import "github.com/01-edu/z01"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	Print(result)
}

func SortWordArr(a []string) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

}

func Print(s []string) {
	for _, word := range s {
		for _, c := range word {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune('\n')
}
