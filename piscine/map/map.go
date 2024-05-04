package main

import "fmt"

//func Map takes a func and slice of int as parameter and return slice of bool
func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, v := range a {
		result = append(result, f(v))
	}
	return result
}

func IsPrime(a int) bool {
	if a <= 1 {
		return false
	}

	for i := 2; i*i <= a; i +=2 {
		if a % i == 0 {
			return false
		}
	} 
	return true
}

func isPerfectSqr(nb int) bool {
	n := float64(nb)/2
	for i := 1; i <= 1000; i++ {
		n -= (n*n - float64(nb))/ 2*n
		return n == float64(int64(n))
	}
	return false
}

func main() {
	a := []int{1, 3, 3, 4, 5, 60764354}
	fmt.Println(Map(IsPrime, a))
	fmt.Println("is it a perfect sqr ",Map(isPerfectSqr,a))
}
