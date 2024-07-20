package main
import (
	"fmt"
)


// function CountIf that returns, the number of elements of a string slice,
// for which the f function returns true
func CountIf(f func(string) bool, tab []string) int {
	count := 0 

	for _, v := range(tab) {
		if f(v) {
			count++
		}
	}
	return count
}
// func IsNumeric takes string and return true if str is numeric
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
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "4", "on", "you", "400"}

	result1 := CountIf(IsNumeric, a1)
	result2 := CountIf(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}