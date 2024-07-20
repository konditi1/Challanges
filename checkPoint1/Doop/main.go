package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	a1 := os.Args[1]
	a2 := os.Args[2]
	a3 := os.Args[3]
	operator := map[string]string{
		"+": "add",
		"-": "sub",
		"*": "mul",
		"/": "div",
		"%": "mod",
	}
	if _, ok := operator[string(a2)]; !ok {
		return
	}

	num, err := Atoi(a1)
	if err != "" {
		os.Stdout.WriteString(err + "\n")
		return
	}

	num2, err := Atoi(a3)
	if err != "" {
		os.Stdout.WriteString(err + "\n")
		return
	}
	result := 0
	switch string(a2) {
	case "+":
		result = num + num2
	case "-":
		result = num - num2
	case "*":
		result = num * num2
	case "/":
		if num2 == 0 {
			os.Stdout.WriteString("No division by 0" + "\n")
			return
		}
		result = num / num2
	case "%":
		if num2 == 0 {
			os.Stdout.WriteString("No modulo by 0" + "\n")
			return
		}
		result = num % num2
	}
	os.Stdout.WriteString(Itoa(result) + "\n")
}

func Atoi(s string) (int, string) {
	var n int
	neg := false
	if s[0] == '-' {
		neg = true
		s = s[1:]
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, "Error converting string to int"
		}
		dig := int(c - '0')
		n = n*10 + dig
	}
	if neg {
		return -n, ""
	}
	return n, ""
}

func Itoa(n int) string {

	if n == 0 {
		return "0"
	}
	str := ""
	neg := false

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		str = string('0'+n%10) + str
		n /= 10
	}
	if neg {
		str = "-" + str
	}
	return str
}
