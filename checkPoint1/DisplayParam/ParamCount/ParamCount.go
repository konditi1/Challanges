package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	count := len(os.Args[1:])
	for _, c := range Itoa(count) {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func Itoa(nb int) string {
	//var str string
	var result []byte
	neg := false
	if nb < 0 {
		nb = -nb
		neg = true
	}
	for nb > 0 {
		digit := (nb % 10) + '0'
		//str = string(digit) + str
		result = append([]byte{byte(digit)}, result...)
		nb = nb / 10
	}

	if neg {
		//str = "-" + str
		result = append([]byte{byte('-')}, result...)
	}
	//return str
	return string(result)
}
