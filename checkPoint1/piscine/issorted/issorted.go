package main

import(
	"fmt"
)

// function IsSorted() that returns true, if the slice of int else false
func IsSorted(f func(a, b int) int, a []int) bool {
	
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == 1 {
			return false
		}
	} 
	return true
}


func main() {
	a1 := []int{0, 1, 2, 3, 8, 5}
	a2 := []int{0, 2, 1, 3}
// takes in a and b  returns a positive int if the first argument is greater
// than the second argument, it returns 0 if they are equal and it returns a negative int otherwise.
	f := func(a, b int)int {
		if a > b {
			return 1
		}
		if a == b {
			return 0
		}
		return -1
	}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
