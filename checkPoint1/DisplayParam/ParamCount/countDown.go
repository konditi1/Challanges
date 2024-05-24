package main

import (
	"fmt"
	"os"
	"strconv"
)

var romanValues = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman, calculation := toRoman(num)
	fmt.Println(calculation)
	fmt.Println(roman)
}

func toRoman(num int) (string, string) {
	var roman, calculation string
	for _, v := range romanValues {
		for num >= v.value {
			roman += v.symbol
			if calculation != "" {
				calculation += "+"
			}
			calculation += v.symbol
			num -= v.value
		}
	}
	return roman, calculation
}
