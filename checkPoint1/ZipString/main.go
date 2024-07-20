package main

import (
	"fmt"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
	fmt.Println(ZipString("aaaa"))
	fmt.Println(ZipString("zzzzzZZZZZZ"))
}

func ZipString(s string) string {

	var result string
	count := 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			result += Itoa(count) + string(s[i])
			count = 1
		}
	}

	result += Itoa(count) + string(s[len(s)-1])

	return result
}

func Itoa(n int) string {
	str := ""
	if n == 0 {
		return "0"
	}
	for n > 0 {
		str = string('0'+(n%10)) + str
		n /= 10
	}
	return str
}
