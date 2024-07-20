package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 2 || IsEmpty(args[1]) {
		fmt.Println("Expecting only two arguments")
		return
	}
	Print(Split(args[1]))
	
}

func Split(s string) string {
	str := ""
	arr := []string{}

	for _, c := range s {
		if string(c) == " " {
			if str != "" {
				arr = append(arr, str)
				str = ""
			}
		} else {
			str += string(c)
		}
	}

	if str != "" {
		arr = append(arr, str)
	}
	return arr[len(arr)-1]
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)

	}
	z01.PrintRune('\n')
}

func IsEmpty(s string) bool {
	for _, c := range s {
		if c != ' ' {
			return false
		}
	}
	return true
}
