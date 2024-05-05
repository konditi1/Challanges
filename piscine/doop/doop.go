package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	fmt.Println("line 12")


	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]
	// arg1 := 9223372036854775807
	// arg2 := "/"
	// arg3 := 0

	var res int

	switch arg2 {
	case "+":
		res = sum(arg1, arg3)

	case "-":
		res = sub(arg1, arg3)

	case "/":
		if arg3 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		res = div(arg1, arg3)

	case "*":
		res = mult(arg1, arg3)

	case "%":
		if arg3 == 0 {
			os.Stdout.WriteString("No modulo by 0")
			return
		}
		res = modulo(arg1, arg3)
	default:
		return
	}

	if res < -9223372036854775807 || res > 9223372036854775807 {
		os.Exit(0)
	}

	fmt.Println(res)
}

func sum(a, b int) int {
	return (a + b)
}

func sub(a, b int) int {
	return (a - b)
}

func div(a, b int) int {
	return (a / b)
}

func mult(a, b int) int {
	return (a * b)
}

func modulo(a, b int) int {
	return (a % b)
}
