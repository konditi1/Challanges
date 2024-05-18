package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main function takes a path to a data file as argument
// and prints the result of each statistic mentioned above
func main() {
	// validate command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run your-program.go <path_to_data_file> ")
		os.Exit(1)
	}

	// get path to data file from command line argument
	filepath := os.Args[1]

	// read data from file
	dataRead := ReadFile(filepath)

	// convert string to float64
	arrData := alphabeToFloat(dataRead)

	// calculate statics
	fmt.Printf("Average: %.0f\n", Average(arrData))
	fmt.Printf("Median: %.0f\n", Median(arrData))
	fmt.Printf("Variance: %.0f\n", Variance(arrData))
	fmt.Printf("Standard deviation: %.0f\n", Sqrt(Variance(arrData)))

}


// func Median takes arr of float64 sorts array of float64
// Returns the median
func Median(arr []float64) float64 {
	// Sort the array
	Sort(arr)	
	
	n := len(arr)
	if n%2 == 0 {
		return (arr[(n/2 - 1)] + arr[n/2]) / 2
	}
	return arr[n/2]
}



// Sort sorts the given array of float64 in ascending order
func Sort(arr []float64) {
	// Bubble sort algorithm, Time complexity: O(n^2)
	// Space complexity: O(1)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				// swap the two elements
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// ReadFile reads the contents of a text file into a slice of strings
// Returns: slice of strings containing the contents of the file
func ReadFile(filepath string) []string {

	// check if file has .txt suffix
	if !strings.HasSuffix(filepath, ".txt") {
		fmt.Println("The file must be a .txt file")
		os.Exit(1)
	}

	// check if file exists
	_, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read file
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err) // print error
		os.Exit(1)                             // exit with error
	}
	// split string
	data1 := strings.Split(string(data), "\n") // split string by newline

	return data1 // return slice of strings
}



// alphabeToFloat takes a slice of strings and converts each element to
// a float64 and Returns: slice of floats
func alphabeToFloat(s []string) []float64 {
	var arrData []float64 

	for i := 0; i < len(s); i++ {
		data, err := strconv.ParseFloat(s[i], 64)
		if err != nil {
			fmt.Println("Error converting string to int :", err)
			continue
		}
		arrData = append(arrData, data)
	}
	return arrData
}



// Average It takes a slice of float64 and returns the average as a float64
func Average(arr []float64) float64 {
	total := 0.0

	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}

	return total / float64(len(arr))
}

// Variance It takes a slice of float64 and returns the variance as a float64
func Variance(arr []float64) float64 {
	mean := Average(arr)
	variance := 0.0
	for i := 0; i < len(arr); i++ {
		variance += (arr[i] - mean) * (arr[i] - mean)
	}
	return variance / float64(len(arr))
}

// Sqrt It takes a float64 and returns the square root of the float64
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
