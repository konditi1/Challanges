package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("args must be 2")
		return
	}

	fmt.Println(AlphaMirror(os.Args[1]))
}

func AlphaMirror(s string) string {

	str := ""

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			str += string('z' - (c -'a'))
		} else if c >= 'A' && c <= 'Z' {
			str += string('Z' - (c -'A'))
		} else {
			str += string(c)
		}
	}
	return str
}