package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	num, err := Atoi(os.Args[1])
	if err != "nil" {
		Print(err)
		return
	}
	calc, rom := InToRoman(num)

	Print(calc)
	z01.PrintRune(' ')
	Print(rom)
	z01.PrintRune('\n')

}

func InToRoman(n int) (string, string) {
	roman := []struct {
		num         int
		roman       string
		calculation string
	}{
		{1000, "M", "M"},
		{900, "CM", "(M-C)"},
		{500, "D", "D"},
		{400, "CD", "(D-C)"},
		{100, "C", "C"},
		{90, "XC", "(C-X)"},
		{50, "L", "L"},
		{40, "XL", "(L-X)"},
		{10, "X", "X"},
		{9, "IX", "(X-I)"},
		{5, "V", "V"},
		{4, "IV", "(V-I)"},
		{1, "I", "I"},
	}

	rom := ""
	calc := ""

	for _, r := range roman {
		for n >= r.num {
			rom += r.roman
			if calc == "" {
				calc += r.calculation
			} else {
				calc += "+" + r.calculation
			}
			n -= r.num
		}
	}

	return calc, rom
}

func Atoi(s string) (int, string) {
	var n int
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, "ERROR: cannot convert to roman digit"
		}
		digit := int(c - '0')
		n = n*10 + digit
	}
	return n, "nil"
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	//z01.PrintRune('\n')
}
