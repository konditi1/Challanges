package main

import (
	"fmt"
)

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s1 string) string {
	s := ToLower(s1)
	str := ""
	for i := 0; i < len(s); i++ {
		if i > 0 && !IsAlpha(string(s[i-1])) && IsAlpha(string(s[i])) {
			if i 
			str += ToUpper(string(s[i]))
		} else {
			if i == 0 {
				str += ToUpper(string(s[i]))
				i++
			}
			str += string(s[i])
		}
	}
	return str
}

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			return true
		}
	}
	return false
}

func ToLower(s string) string {
	var str string
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			str += string(s[i] + 32)
		} else {
			str += string(s[i])
		}
	}
	return str
}

func ToUpper(s string) string {
	return string(s[0] - 32)
}
