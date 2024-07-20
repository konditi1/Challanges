package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(9))
}

func FindPrevPrime(nb int) int {

	for nb > 0 {
		if IsPrime(nb) {
			return nb
		}
		nb--
	}
	return 0
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} 

	if nb == 2 || nb == 3 {
		return true
	}

	if nb % 2 == 0 || nb % 3 == 0 {
		return false
	}

	for i := 5; i*i <= nb; i += 6 {
		if nb % i == 0 || nb % (i + 2) == 0 {
			return false
		}
	}
	return true
}
