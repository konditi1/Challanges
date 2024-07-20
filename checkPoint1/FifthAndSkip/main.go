package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	
	if len(str) == 0 {
		return "\n"
	}

	s1 := strings.Split(str, " ")
	s2 := strings.Join(s1, "")
	s3 := ""
	if len(s2) < 5 {
		return "Invalid Input\n"
	}
	for i := 0; i < len(s2); i += 6 {
		end := 5 + i
		if end > len(s2) {
			end = len(s2)
		}
		if len(s3) != 0 {
			s3 += " "
		} 
			s3 +=  string(s2[i : end])
		

	}
	

	return s3 + "\n"
}
