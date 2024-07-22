package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"
	str := ""
	word := ""
	a := []string{}

	for _, n := range arr {
		if n < ' ' || n > '~' {
			word += "."
		} else {
			word += string(n)
		}
		for n > 0 {
			str = string(hex[n%16]) + str
			n /= 16
		}
		a = append(a, str)
		str = ""
	}

	count := 0
	for i, num := range a {
		count++
		if num == "" {
			num = "00"
		}
		for _, r := range num {
			z01.PrintRune(r)
		}
		if count%4 == 0 {
			z01.PrintRune(10)
		}else if i != len(a)-1 {
			z01.PrintRune(' ')
		} 
	}
	z01.PrintRune(10)
	Print(word)
}

func PrintMemory1(arr [10]byte) {
	var a []string
	str := ""
	str1 := ""
	hex := "0123456789abcdef"

	for _, v := range arr {
		if v < 32 || v > 126 {
			str1 += "."
		} else {
			str1 += string(v)
		}
		for v > 0 {
			rem := v % 16
			str = string(hex[rem]) + str
			v = v / 16
		}
		a = append(a, str)
		str = ""
	}

	count := 0

	for i := 0; i <= len(a)-1; i++ {
		count++
		if a[i] == "" {
			a[i] = "00"
		}
		for _, v := range a[i] {
			z01.PrintRune(v)
		}
		if count%4 == 0 {
			z01.PrintRune(10)
		} else if i == len(a)-1 {
			break
		} else {
			z01.PrintRune(' ')

		}
	}
	z01.PrintRune(10)

	Print(str1)

}

func Print(s string) {
	for _, val := range s {
		z01.PrintRune((val))
	}
	z01.PrintRune(10)
}
