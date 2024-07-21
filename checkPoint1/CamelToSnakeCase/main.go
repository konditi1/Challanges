package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase(" "))
	fmt.Println(CamelToSnakeCase(""))
	fmt.Println(CamelToSnakeCase("A"))
	fmt.Println(CamelToSnakeCase("snack"))
	fmt.Println(CamelToSnakeCase("thisIsaTestOfCamelCase"))
	fmt.Println(CamelToSnakeCase("aA"))
}

func CamelToSnakeCase(s string) string {
	var SnakeCase string
	if len(s) == 0 {
		return ""
	}
	if IsCaps(s[len(s)-1]) {
		return s
	}
	for i := range s {
		if i+1 < len(s) && IsCaps(s[i]) && IsCaps(s[i+1]) {
			return s
		} else if IsDigit(s[i]) {
			return s
		}
		if IsCaps(s[i]) {
			if i != 0 {
				SnakeCase += "_"
			}
			SnakeCase += string(s[i])
		} else {
			SnakeCase += string(s[i])
		}
	}
	return SnakeCase
}

func IsDigit(n byte) bool {
	if n > '0' && n < '9' {
		return true
	}
	return false
}

func IsCaps(s byte) bool {
	if s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}
