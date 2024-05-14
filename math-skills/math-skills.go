package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run your-program.go <path_to_data_file> ")
		os.Exit(1)
	}

	filepath := os.Args[1]

	dataRead := ReadFile(filepath)

	arrData := alphabeToFloat(dataRead)

	fmt.Printf("Average: %.0f\n", Average(arrData))
	fmt.Printf("Median: %.0f\n", Median(arrData))
	fmt.Printf("Variance: %.0f\n", Variance(arrData))
	fmt.Printf("Standard deviation: %.0f\n", Sqrt(Variance(arrData)))
	
}

func Median(arr []float64) float64 {
	Sort(arr)
	n := len(arr)
	
	if n % 2 == 0 {
		return (arr[(n/2 - 1)] + arr[n/2]) /2
	}
	return arr[(n/2)]
}

func Sort(arr []float64) {
	for i := 0; i < len(arr); i++ {
		for j := i +1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func  ReadFile(filepath string) []string {

	// check if file is ,txt
	if !strings.HasSuffix(filepath, ".txt") {
		fmt.Println("The file must be .txt")
		os.Exit(1)
	}

	//check if file exists
	_, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read file
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file :", err)
		os.Exit(1)
	}
	// split string
	data1 := strings.Split(string(data), "\n")

	return data1
}


func alphabeToFloat(s []string) []float64 {
	var arrData []float64
	for i := 0; i < len(s); i++ {
		// convert string to int
		data, err := strconv.ParseFloat(s[i], 64)
		if err != nil {
			fmt.Println("Error converting string to int :", err)
			continue
		}
		arrData = append(arrData, data)
	}
	return arrData
}


func Average(arr []float64) float64 {
	total := 0.0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}
	return total / float64(len(arr))
}

func Variance(arr []float64) float64 {
	mean := Average(arr)
	variance := 0.0
	for i := 0; i < len(arr); i++ {
		variance += (arr[i] - mean) * (arr[i] - mean)
	}
	return variance / float64(len(arr))
}

func Sqrt(nb float64) float64 {
	s := nb/2
	prev := s
	for i := 0; i < 10000; i++ {
		s -= (s*s - nb) / (2 * s)
		if s == prev {
			break
		}
		prev = s
	}
	return s
}


// func Atoi takes a string and returns an int and an error
// func Atoi(s string) (int, error) {
// 	nb := 0
// 	sign := 1
// 	if len(s) > 0 && s[0] == '-' {
// 		sign = -1
// 	}
// 	for i := len(s) - 1; i >= 0; i-- {
// 		v := s[i]
// 		if v < '0' || v > '9' {
// 			return 0, fmt.Errorf("non integer character")
// 		}
// 		nb = nb*10 + int(v-'0')
// 	}
// 	return nb * sign, nil
// } // convert string to int

// func Split(s string, sep string) []string {
// 	var result []string
// 	start := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i:i+len(sep)] == sep {
// 			result = append(result, s[start:i])
// 			start = i + len(sep)
// 			i += len(sep) - 1
// 		}
// 		result = append(result, s[start:])
// 	}
// 	return result
// }
