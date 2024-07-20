package main

import "fmt"


// function Any takes a function literal and slice the returns true if any condtion is true
func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}

func IsNumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, c := range str {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func main() {
	a1 := []string{"Hello", "how", "44", "you"}
	a2 := []string{"This", "i4s", "on", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
