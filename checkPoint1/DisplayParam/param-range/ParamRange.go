package main

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	numArr := []int{}

	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	for _, arg := range args {
		num, err := Atoi(arg)
		if err != nil {
			PrintInRunes("Error")
			z01.PrintRune('\n')
			return
		}
		numArr = append(numArr, num)
	}
	fmt.Println(numArr)

	sort.Ints(numArr)

	fmt.Println(numArr[0], numArr[len(numArr)-1])
}

func Atoi(s string) (int, error) {
	var nb int
	neg := false

	if len(s) == 0 {
		return 0, errors.New("empty string input")
	}

	if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	for  i := 0; i <= len(s) - 1; i++ {		
		if s[i] < '0' || s[i] > '9' {
			return 0, errors.New("invalid character in input: " + string(s[i]))
		}
		//nb = int(s[i] - '0')
		nb = int(s[i]-'0') + nb * 10 
	}
	if neg {
		nb = -nb
	}
	return nb, nil
}

func PrintInRunes(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
