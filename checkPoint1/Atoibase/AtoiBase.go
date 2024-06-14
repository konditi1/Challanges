package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("1111101", "01"))          // Binary to Decimal
	fmt.Println(AtoiBase("125", "0123456789"))      // Decimal to Decimal (should be 125)
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF")) // Hexadecimal to Decimal
	fmt.Println(AtoiBase("uoi", "choumi"))          // Custom base "choumi" to Decimal
	fmt.Println(AtoiBase("bbbbbab", "-ab"))         // Custom base with '-' (if needed)
}

func AtoiBase(str string, base string) int {
	var baseMap = make(map[rune]int)
	for i, c := range base {
		if c == '-' || c == '+' {
			return 0
		}
		baseMap[c] = i
	}

	power := 1
	result := 0
	for i := len(str) - 1; i >= 0; i-- {
		val, ok := baseMap[rune(str[i])]
		if !ok {
			return -1
		}
		result += power * val
		power *= len(base)
	}
	return result
}
