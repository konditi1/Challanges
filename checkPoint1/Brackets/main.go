package main

import (
	"fmt"
	"os"
)


func Brackets(s string) string {
	
	closingBracketSlice := []rune{}
	
	brackets := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, c := range s {
		
		if closing, ok := brackets[c]; ok {
			closingBracketSlice = append(closingBracketSlice, closing)
		} else if c == ')' || c == ']' || c == '}' {
			
			if len(closingBracketSlice) == 0 || closingBracketSlice[len(closingBracketSlice)-1] != c {
				return "Error"
			}
			// Pop the last bracket
			closingBracketSlice = closingBracketSlice[:len(closingBracketSlice)-1]
		}
	}
	
	if len(closingBracketSlice) == 0 {
		return "OK"
	}
	return "Error"
}

func main() {
	args := os.Args[1:] 
	if len(args) == 0 {
		return 
	}

	for _, arg := range args {
		fmt.Println(Brackets(arg))
	}
}
