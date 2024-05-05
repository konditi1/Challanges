package main

import (
	"fmt"
)

func main() {

	// if len(os.Args) != 4 {
	// 	return
	// }

	// arg1 := os.Args[1]
	// arg2 := os.Args[2]
	// arg3 := os.Args[3]
	arg1 := 9223372036854775807
	arg2 := "+"
	arg3 := 1

	res := 

	switch arg2 {
	case "+":
		sum(arg1, arg3)

	case "-":
		sub(arg1, arg3)

	case "/":
		div(arg1, arg3)

	case "*":
		mult(arg1, arg3)

	case "%":
		modulo(arg1, arg3)
	default:
		return
	}

	fmt.Println(arg1, arg2, arg3)

}

func sum(a, b int) interface{} {
	s := a + b
	if a != 0 && s-b != a {
		return nil
	}
	return (a + b)
}

func sub(a, b int) int {
	return (a - b)
}

func div(a, b int) interface{} {
	if b == 0 {
		return "No division by 0"
	}
	return (a / b)
}
func mult(a, b int) int {
	return (a * b)
}

func modulo(a, b int) interface{} {
	if b == 0 {
		return "No modulo by 0"
	}
	return (a % b)
}
