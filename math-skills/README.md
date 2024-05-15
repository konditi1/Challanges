# Statistical Calculations Program
# Objectives
## Overview
```
This project aims to calculate key statistical metrics—Average, Median, Variance, and Standard Deviation—from a dataset provided in a text file. The dataset is read from a file whose path is provided as a command-line argument. Each line in the file represents one value in the statistical population.
```
## Objectives
```
The program calculates the following statistics:
- Average
- Median
- Variance
- Standard Deviation
```
## Instructions

### Prerequisites
```
Ensure you have [Go](https://golang.org/dl/) installed on your machine.
```
### Usage

To run the program, use the following command, replacing `<path_to_data_file>` with the actual path to your data file:

```sh
go run your-program.go <path_to_data_file>

```
### Data File Format
```
Your program must be able to read from a file and print the result of each statistic mentioned above. In other words, your program must be able to read the data present in the path passed as argument. The data in the file will be presented as the following example:

189
113
121
114
145
110
...
This data represents a statistical population: each line contains one value.
```
## Output

```
The program reads the file, computes the required statistics, and prints the results as rounded integers in the following format:

Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
Please note that the values are rounded integers.
```
# Testing
```
The program will be tested by an auditor using an automated testing script. The script generates a random dataset and compares the output of your program against expected results.

You can choose the language in which you want to build your program.
```
## Testing Steps
Prepare a data file named data.txt with random numerical values, one per line.
Run the program using the command provided above.
Verify that the output matches the expected statistical calculations.

## Learning Outcomes
This project helps you understand and apply the following concepts:

Basic statistics and mathematical calculations
Working with files and command-line arguments in Go
Implementing sorting algorithms and basic numerical methods
