package main

import (
    "fmt"
    "math"
    "math/rand"
    "os"
    "sort"
)

func main() {
    // Generate random dataset of numbers
    const (
        datasetSize = 100
        min         = 1
        max         = 1000
        fileName    = "dataset.txt"
    )

    // Create and open file for writing
    file, err := os.Create(fileName)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Generate random numbers and write to file
    for i := 0; i < datasetSize; i++ {
        num := rand.Intn(max-min+1) + min
        _, err := fmt.Fprintf(file, "%d\n", num)
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    }

    // Close the file
    if err := file.Close(); err != nil {
        fmt.Println("Error closing file:", err)
        return
    }

    // Open the file for reading
    file, err = os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Read the numbers from the file into a slice
    var numbers []int
    var num int
    for {
        _, err := fmt.Fscanf(file, "%d\n", &num)
        if err != nil {
            break
        }
        numbers = append(numbers, num)
    }

    // Calculate statistics
    average := calculateAverage(numbers)
    median := calculateMedian(numbers)
    variance := calculateVariance(numbers, average)
    stdDev := math.Sqrt(variance)

    // Print results
    fmt.Println("Average:", average)
    fmt.Println("Median:", median)
    fmt.Println("Variance:", variance)
    fmt.Println("Standard Deviation:", stdDev)
}

func calculateAverage(numbers []int) float64 {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return float64(sum) / float64(len(numbers))
}

func calculateMedian(numbers []int) float64 {
    sort.Ints(numbers)
    n := len(numbers)
    if n%2 == 0 {
        return float64(numbers[n/2-1]+numbers[n/2]) / 2
    }
    return float64(numbers[n/2])
}

func calculateVariance(numbers []int, mean float64) float64 {
    sum := 0.0
    for _, num := range numbers {
        deviation := float64(num) - mean
        sum += deviation * deviation
    }
    return sum / float64(len(numbers))
}
