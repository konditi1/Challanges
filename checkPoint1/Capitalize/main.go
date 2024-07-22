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
		if i > 0 && !IsAlphaNum(string(s[i-1])) && IsAlpha(string(s[i])) {
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

func IsAlphaNum(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') ||
		(s[i] >= 'A' && s[i] <= 'Z') ||
		(s[i] >= '0' && s[i] <= '9') {
			return true
		}
	}
	return false
}