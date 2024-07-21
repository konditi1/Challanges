package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return a
	}
	if nbrs[0] < 0 {
		nbrs[0] += len(a)
	}

	fmt.Println(nbrs)
	if len(nbrs) == 1 {
		return a[nbrs[0]:]
	}
	if len(nbrs) == 2 {
		if nbrs[1] < 0 {
			nbrs[1] += len(a)
		}
		if nbrs[0] > nbrs[1] {
			return nil
		}
		return a[nbrs[0]:nbrs[1]]
	}
	return a
}
