### LEVEL 2
## countdown

### Instructions

Write a program that displays all digits in descending order, followed by a newline (`'\n'`).

### Usage

```console
$ go run .
9876543210
$
```
```
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := '9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
```

## max

### Instructions

Write a function `Max` that will return the maximum value in a slice of integers. If the slice is empty it will return 0.

### Expected function

```go
func Max(a []int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(a)
	fmt.Println(max)
}
```

And its output :

```console
$ go run .
123
$
```
```
func Max(a []int)int {
	if len(a) == 0 {
		return 0
	}
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return(a[len(a)-1])
}
```

## strlen

### Instructions

- Write a function that counts the `runes` of a `string` and that returns that count.

### Expected function

```go
func StrLen(s string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
```

And its output :

```console
$ go run .
12
$
```
```
func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
```

## firstrune

### Instructions

Write a function that returns the first `rune` of a `string`.

### Expected function

```go
func FirstRune(s string) rune {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}
```

And its output :

```console
$ go run .
HSO
$
```

### Notions

- [rune-literals](https://golang.org/ref/spec#Rune_literals)

```
func FirstRune(s string) rune {

	return rune(s[0])
}
```
## lastrune

### Instructions

Write a function that returns the last `rune` of a `string`.

### Expected function

```go
func LastRune(s string) rune {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}
```

And its output :

```console
$ go run .
!!!
$
```

### Notions

- [rune-literals](https://golang.org/ref/spec#Rune_literals)

```
func LastRune(s string) rune {
	return rune(s[len(s)-1])
}
```
## lastword

### Instructions

Write a program that takes a `string` and displays its last word, followed by a newline (`'\n'`).

- A word is a section of `string` delimited by spaces or by the start/end of the `string`.

- The output will be followed by a newline (`'\n'`).

- If the number of arguments is different from 1, or if there are no word, the program displays nothing.

### Usage

```console
$ go run . "FOR PONY" | cat -e
PONY$
$ go run . "this        ...       is sparta, then again, maybe    not" | cat -e
not$
$ go run . "  "
$ go run . "a" "b"
$ go run . "  lorem,ipsum  " | cat -e
lorem,ipsum$
$
```
```
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 2 || len(args[1]) == 0 || IsEmpty(args[1]) {
		return
	}

	for _, c := range Split(args[1]) {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')

}

func Split(s string) string {
	var str string
	var arr []string

	for _, c := range s {
		if string(c) == " " {
			if str != "" {
				arr = append(arr, str)
				str = ""
			}
		} else {
			str += string(c)
		}
	}
	if len(str) != 0 {
		arr = append(arr, str)
	}

	return arr[len(arr)-1]
}

func IsEmpty(s string) bool {
	for _, c := range s {
		if string(c) != " " {
			return false
		}
	}
	return true
}
```

## reduceint

### Instructions

The function should have as parameters a slice of integers `a []int` and a function `f func(int, int) int`.  For each element of the slice, it should apply the function `f func(int, int) int`, save the result and then print it.

### Expected function

```go
func ReduceInt(a []int, f func(int, int) int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

```

And its output :

```console
$ go run .
1000
502
250
$
```
```
func ReduceInt(a []int, f func(int, int) int) {
	reduced := a[0]
	for i := 1; i < len(a); i++ {
		reduced = f(reduced, a[i])
	}
	PrintNumRune(reduced)
	z01.PrintRune('\n')
}

func PrintNumRune(n int) {
	if n == 0 {
		return
	}
	PrintNumRune(n/10)
	z01.PrintRune(rune('0' + (n % 10)))
}

```

## rot13 

### Instructions

Write a program that takes a `string` and displays it, replacing each of its
letters by the letter 13 spots ahead in alphabetical order.

- 'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

- The output will be followed by a newline (`'\n'`).

- If the number of arguments is different from 1, the program displays nothing.

### Usage

```console
$ go run . "abc"
nop
$ go run . "hello there"
uryyb gurer
$ go run . "HELLO, HELP"
URYYB, URYC
$ go run .
$
```
```
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		return
	}
	Rot13(os.Args[1])
	z01.PrintRune('\n')
}

func Rot13(s string) {
	for _, c := range os.Args[1] {
		if c >= 'a' && c <= 'm' {
			z01.PrintRune(c + 13)
		} else if c >= 'n' && c <= 'z' {
			z01.PrintRune(c - 13)
		} else if c >= 'A' && c <= 'M' {
			z01.PrintRune(c + 13)
		} else if c >= 'N' && c <= 'Z' {
			z01.PrintRune(c - 13)
		} else {
			z01.PrintRune(c)
		}
	}
}
```

## rot14

### Instructions

Write a function `rot14` that returns the `string` within the parameter transformed into a `rot14 string`.
Each letter will be replaced by the letter 14 spots ahead in the alphabetical order.

- 'z' becomes 'n' and 'Z' becomes 'N'. The case of the letter stays the same.

### Expected function

```go
func Rot14(s string) string {

}
```

### Usage

Here is a possible program to test your function :

```
package main

import (
	"piscine"
	"github.com/01-edu/z01"
)

func main() {
	result := piscine.Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
```

And its output :

```console
$ go run .
Vszzc! Vck ofs Mci?
$
```

```
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Rot14(s string) string {
	var str string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			str += string((c-'a'+14)%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			str += string((c-'A'+14)%26 + 'A')
		} else {
			str += string(c)
		}
	}
	return str
}
```
## alphamirror

### Instructions

Write a program called `alphamirror` that takes a `string` as argument and displays this `string` after replacing each alphabetical character with the opposite alphabetical character.

The case of the letter remains unchanged, for example :

'a' becomes 'z', 'Z' becomes 'A'
'd' becomes 'w', 'M' becomes 'N'

The final result will be followed by a newline (`'\n'`).

If the number of arguments is different from 1, the program prints a new line.

### Usage

```console
$ go run . "abc"
zyx$
$
$ go run . "My horse is Amazing." | cat -e
Nb slihv rh Znzarmt.$
$
$ go run .
$
```
```
package main

import(
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	PrintRune(AlphaMirror(os.Args[1]))
	z01.PrintRune('\n')
}

func AlphaMirror(s string) string {
	var str string

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			str += string('z' - (c - 'a'))

		} else if c >= 'A' && c <= 'Z' {
			str += string('Z' - (c - 'A'))
		} else {
			str += string(c)
		}
	}
	return str
}

func PrintRune(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
```

## chunk

### Instructions

Write a function called `Chunk` that receives as parameters a slice, `slice []int`, and a number `size int`. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of `size`.

- If the `size` is `0` it should print a newline (`'\n'`).

### Expected function

```go
func Chunk(slice []int, size int) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
```

And its output :

```console
$ go run .
[]

[[0 1 2] [3 4 5] [6 7]]
[[0 1 2 3 4] [5 6 7]]
[[0 1 2 3] [4 5 6 7]]
$
```
```
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 99}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 99}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {

	var result [][]int
	if size == 0 {
		z01.PrintRune('\n')
		return
	}

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	z01.PrintRune('[')
	for i, arr := range result {
		z01.PrintRune('[')
		for j, num := range arr {
			if num == 0 {
				z01.PrintRune('0')
			}
			PrintNumInRunes(num)
			if j < len(arr)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')
		if i < len(result)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}

// func PrintNumInRunes(n int) {
// 	if n == 0 {
// 		return
// 	}
// 	PrintNumInRunes(n/10)
// 	z01.PrintRune (rune('0' + n % 10))
// }

func PrintNumInRunes(n int) {
	var str string
	if n == 0 {
		str += string(n)
	} else {
		for n > 0 {
			str += string('0' + n % 10)
			n = n/10
		}
	}
	for _, c := range str {
		z01.PrintRune(c)
	}
}

```
## compare

### Instructions

Write a function that behaves like the `Compare` function.

### Expected function

```go
func Compare(a, b string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "lut!"))
	fmt.Println(piscine.Compare("Ola!", "Ol"))
}
```

And its output :

```console
$ go run .
0
-1
1
$
```

### Notions

- [strings/Compare](https://golang.org/pkg/strings/#Compare)
```
func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
```
## nextprime

### Instructions

Create a **function** which returns the first prime number which is greater than or equal to the `u64` passed as an argument.

The function must be optimized, so as to avoid time-outs.

> We consider that only positive numbers can be prime numbers.

### Expected function

```rust
pub fn next_prime(nbr: u64) -> u64 {

}
```

### Usage

Here is a possible program to test your function :

```rust
use nextprime::*;

fn main() {
    println!("The next prime after 4 is: {}", next_prime(4));
    println!("The next prime after 11 is: {}", next_prime(11));
}

```

And its output :

```console
$ cargo run
The next prime after 4 is: 5
The next prime after 11 is: 11
$
```
## findprevprime

### Instructions

Write a function that returns the first prime number that is equal or inferior to the `int` passed as parameter.

If there are no primes inferior to the `int` passed as parameter the function should return 0.

### Expected function

```go
func FindPrevPrime(nb int) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
```

And its output :

```console
$ go run .
5
3
$
```
```
func FindPrevPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	}
	for i := nb; i > 0; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}

func IsPrime(n int) bool {
	if n == 0 || n ==1 {
		return false
	}

	if n == 2 || n == 3{
		return true
	}
	
	if n % 2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i +=2 {
		if n % i == 0 { 
			return false
		}
	}
	return true
}
```

## foldint

### Instructions

The function should have as parameters a function, `f func(int, int) int` a slice of integers, `slice []int` and an int `acc int`. For each element of the slice, it should apply the arithmetic function, save the result and print it. The function will be tested with our own functions `Add, Sub, and Mul`.

### Expected function

```go
func FoldInt(f func(int, int) int, a []int, n int) {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

```

And its output :

```console
$ go run .
99
558
87

93
0
93
$
```
```
func FoldInt(f func(int, int) int, a []int, n int) {

	result := n

	for _, num := range a {
		result = f(result, num)
	}
	if result == 0 {
		z01.PrintRune('0')
	}
	PrintNumRunes(result)
	z01.PrintRune('\n')

}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func PrintNumRunes(n int) {
	if n == 0 {
		return
	}
	PrintNumRunes(n / 10)
	z01.PrintRune(rune('0' + n%10))
}
```
## searchreplace

### Instructions

Write a program that takes 3 arguments, the first argument is a `string` in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

- If the number of arguments is different from 3, the program displays nothing.

- If the second argument is not contained in the first one (the string) then the program rewrites the `string` followed by a newline (`'\n'`).

### Usage

```console
$ go run . "hella there" "a" "o"
hello there
$ go run . "hallo thara" "a" "e"
hello there
$ go run . "abcd" "z" "l"
abcd
$ go run . "something" "a" "o" "b" "c"
$
```
```
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}
	PrintStr(SearchReplace(args[1], args[2], args[3]))
	z01.PrintRune('\n')
}

func SearchReplace(s, letter, repleca string) string {
	var str string 
	for _, c := range s {
		if string(c) == letter {
			str += repleca
		} else {
			str += string(c)
		}
	}
	return str
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
```
## tabmult

### Instructions

Write a program that displays a number's multiplication table.

- The parameter will always be a strictly positive number that fits in an `int`. Said number multiplied by 9 will also fit in an `int`.

### Usage

```console
$ go run . 9
1 x 9 = 9
2 x 9 = 18
3 x 9 = 27
4 x 9 = 36
5 x 9 = 45
6 x 9 = 54
7 x 9 = 63
8 x 9 = 72
9 x 9 = 81
$ go run . 19
1 x 19 = 19
2 x 19 = 38
3 x 19 = 57
4 x 19 = 76
5 x 19 = 95
6 x 19 = 114
7 x 19 = 133
8 x 19 = 152
9 x 19 = 171
$ go run .

$
```
```
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		return
	}
	TabMult(args[1])
}

func TabMult(n string) {
	str1 := " x "
	str2 := " = "
	num := Atoi(n)
	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune('0' + i))
		PrintStr(str1)
		PrintStr(n)
		PrintStr(str2)
		PrintNumInRunes(i * num)
		z01.PrintRune('\n')

	}
}

func Atoi(s string) int {
	var n int
	for _, c := range s {
		digit := int(c - '0')
		n = n*10 + digit

	}

	return n
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func PrintNumInRunes(nb int) {
	var s string
	neg := false
	if nb == 0 {
		z01.PrintRune('0')
		return
	}
	if nb < 0 {
		nb = -nb
		neg = true
	}
	for nb > 0 {
		s = string(nb%10+'0') + s
		nb = nb / 10

	}
	if neg {
		s = "-" + s
	}

	PrintStr(s)
}
```
## inter

### Instructions

Write a program that takes two `string` and displays, without doubles, the characters that appear in both `string`, in the order they appear in the first one.

- The display will be followed by a newline (`'\n'`).

- If the number of arguments is different from 2, the program displays nothing.

### Usage

```console
$ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
$ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
$
```
```
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	str := ""
	// exist := map[rune]bool{}
	// for _, c1 := range s1 {
	// 	if !exist[c1] && IsPresent(s2, c1) {
	// 		str += string(c1)
	// 		exist[c1] = true
	// 	}
	// }
	for _, c1 := range s1 {
		found := false
		for _, c2 := range s2 {
			if c1 == c2 {
				found = true
				//break
			}
		}
		if found && !IsPresent(str, c1) {
			str += string(c1)
		}
	}
	Print(str)
	z01.PrintRune('\n')
}

func IsPresent(s string, val rune) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
```
## ispowerof2

### Instructions

Write a program that determines if a given number is a power of 2. A number `n` is a power of 2 if it meets the following condition: **n = 2 <sup>m</sup>** where m is another integer number.

For example, **4 = 2 <sup>2</sup>** or **16 = 2 <sup>4</sup>** are power of 2 numbers while 24 is not.

This program must print `true` if the given number is a power of 2, otherwise it has to print `false`.

- If there is more than one or no argument, the program should print nothing.

- When there is only one argument, it will always be a positive valid `int`.

### Usage :

```console
$ go run . 2 | cat -e
true$
$ go run . 64
true
$ go run . 513
false
$ go run .
$ go run . 64 1024
$
```
```
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	n2 := 2
	n1 := os.Args[1]
	IsPower(Atoi(n1), n2)
	z01.PrintRune('\n')

}

func IsPower(n1, n2 int) {
	if n1 <= 0 || n2 <= 0 {
		Print("false")
		return
	}

	if n1 == 1 {
		Print("true")
		return
	}
	
	if n2 == 1 {
		Print("false")
		return
	}

	for n1 > 1 {
		if n1 % n2 != 0 {
			Print("false")
			return
		}
		n1 = n1 /n2
	}
	Print("true")
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func Atoi(s string) int {
	var n int
	for _, num := range s {
		digit := int(num - '0')
		n = n * 10 + digit
	}
	return n
}
```

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
### piglatin
```
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	s := os.Args[1]
	vowel := map[string]bool{"a": true, "e": true, "i": true, "o": true, "u": true,
		"A": true, "E": true, "I": true, "O": true, "U": true}
	hasvowel := false
	for _, c := range s {
		if vowel[string(c)] {
			hasvowel = true
			break
		}
	}
	if !hasvowel {
		Print("No vowels")
		z01.PrintRune('\n')
		return
	}
	for i, c := range s {
		if vowel[string(c)] {
			if i == 0 {
				Print(s + "ay")
				z01.PrintRune('\n')
				return
			} else {
				Print(s[i:] + s[:i] + "ay")
				z01.PrintRune('\n')
				return
			}
		}
	}
	z01.PrintRune('\n')
}

func Print(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
```