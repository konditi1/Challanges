## atoi

### Instructions

- Write a function that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number represented as a `string` in a number represented as an `int`.

- `Atoi` returns `0` if the `string` is not considered as a valid number. For this exercise **non-valid `string` chains will be tested**. Some will contain non-digits characters.

- For this exercise the handling of the signs `+` or `-` **does have** to be taken into account.

- This function will **only** have to return the `int`. For this exercise the `error` result of `Atoi` is not required.

### Expected function

```go
func Atoi(s string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
```

And its output :

```console
$ go run .
12345
12345
0
0
1234
-1234
0
0
$
```

### Notions

- [strconv/Atoi](https://golang.org/pkg/strconv/#Atoi)
17
subjects/atoi/main.go
Unescape
View File
@ -0,0 +1,17 @@
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Atoi("12345"))
	fmt.Println(piscine.Atoi("0000000012345"))
	fmt.Println(piscine.Atoi("012 345"))
	fmt.Println(piscine.Atoi("Hello World!"))
	fmt.Println(piscine.Atoi("+1234"))
	fmt.Println(piscine.Atoi("-1234"))
	fmt.Println(piscine.Atoi("++1234"))
	fmt.Println(piscine.Atoi("--1234"))
}
62
subjects/atoibase-exam/README.md
Unescape
View File
@ -1,62 +0,0 @@

## binarycheck
### Instructions
Write a function that takes an int as an argument and returns 0 if the number is odd, and 1 if the number is even.

Expected function
func BinaryCheck(nbr int32) int {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.BinaryCheck(5))
    fmt.Println(piscine.BinaryCheck(0))
    fmt.Println(piscine.BinaryCheck(8))
    fmt.Println(piscine.BinaryCheck(-9))
    fmt.Println(piscine.BinaryCheck(-4))
}
`
And its output:
`
$ go run .
0
1
1
0
1
`

## byebyefirst

### Instructions

Write a function that takes a slice of `string`'s and returns a new slice without the first element.

- If the slice is empty, return the empty slice.

### Expected function

```go
func ByeByeFirst(strings []string) []string {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
}
```

And its output:

```console
$ go run . | cat -e
[]$
[]$
[second]$
[abcd efg]$
```

## cameltosnakecase

### Instructions

Write a function that converts a `string` from `camelCase` to `snake_case`.

- If the `string` is empty, return an empty `string`.
- If the `string` is not `camelCase`, return the `string` unchanged.
- If the `string` is `camelCase`, return the `snake_case` version of the `string`.

For this exercise you need to know that `camelCase` has two different writing alternatives that will be accepted:

- lowerCamelCase
- UpperCamelCase

Rules for writing in `camelCase`:

- The word does not end on a capitalized letter (CamelCasE).
- No two capitalized letters shall follow directly each other (CamelCAse).
- Numbers or punctuation are not allowed in the word anywhere (camelCase1).

### Expected function

```go
func CamelToSnakeCase(s string) string{

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
```

And its output:

```console
$ go run .
Hello_World
hello_World
camel_Case
CAMELtoSnackCASE
camel_To_Snake_Case
hey2
```
## capitalize

### Instructions

Write a function that capitalizes the first letter of each word **and** lowercases the rest.

A word is a sequence of **alphanumeric** characters.

### Expected function

```go
func Capitalize(s string) string {

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
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
}
```

And its output :

```console
$ go run .
Hello! How Are You? How+Are+Things+4you?
$
```

## canyoucount

### Instructions

Your program will receive some arguments. Count how many characters they have in total and print them.

- If the number of arguments is invalid it should print `0`.

### Usage

```console
$ go run . "hello" "how are you?" | cat -e
17$
$ go run . "hi" | cat -e
2$
$ go run . | cat -e
0$
```

## capitalizing

### Instructions

Complete the `capitalize_first` **function** which converts the first letter of the string to uppercase.

Complete the `title_case` **function** which converts the first letter of each word in a string to uppercase.

Complete the `change_case` **function** which converts the uppercase letters of a string into lowercase, and the lowercase to uppercase.

### Expected Functions

```rust
pub fn capitalize_first(input: &str) -> String {
}

pub fn title_case(input: &str) -> String {
}

pub fn change_case(input: &str) -> String {
}
```

### Usage

Here is a program to test your functions.

```rust
use capitalizing::*;

fn main() {
    println!("{}", capitalize_first("joe is missing"));
    println!("{}", title_case("jill is leaving A"));
    println!("{}",change_case("heLLo THere"));
}
```

And its output

```consoole
$ cargo run
Joe is missing
Jill Is Leaving A
HEllO thERE
$
```


## cleanstr

### Instructions

Write a **program** that takes a `string`, and displays this `string` with exactly:

- one space between words.
- without spaces nor tabs at the beginning nor at the end.
- with the result followed by a newline ("`\n`").

A "word" is defined as a part of a `string` delimited either by spaces/tabs, or
by the start/end of the `string`.

If the number of arguments is not 1, or if there are no words to display, the
program displays a newline("`\n`").

### Usage :

```console
$ go run . "you see it's easy to display the same thing" | cat -e
you see it's easy to display the same thing$
$ go run . " only    it's  harder   "
only it's harder$
$ go run . " how funny" "Did you   hear Mathilde ?"

$ go run . ""

$
```




## atoibase

### Instructions

Write a function that takes two arguments:

- `s`: a numeric `string` in a given [base](<https://simple.wikipedia.org/wiki/Base_(mathematics)>).
- `base`: a `string` representing all the different digits that can represent a numeric value.

And return the integer value of `s` in the given `base`.

If the base is not valid it returns `0`.

Validity rules for a base :

- A base must contain at least 2 characters.
- Each character of a base must be unique.
- A base should not contain `+` or `-` characters.

String number must contain only elements that are in base.

Only valid `string` numbers will be tested.

The function **does not have** to manage negative numbers.

### Expected function

```go
func AtoiBase(s string, base string) int {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
```

And its output :

```console
$ go run .
125
125
125
125
0
$
```
14
subjects/atoibase/main.go
Unescape
View File
@ -0,0 +1,14 @@
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))
	fmt.Println(piscine.AtoiBase("1111101", "01"))
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))
}
11
## bezero

### Instructions
- Write a function that takes a slice of integers as an argument and returns the same slice by initializing all the elements to 0.

If the slice is empty, return an empty slice.
### Expected function
`
func BeZero(slice []int) []int {

}
`
subjects/bezero/main.go
Unescape
View File
@ -0,0 +1,11 @@
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BeZero([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(piscine.BeZero([]int{}))
}
`
$ go run .
[0 0 0 0 0 0]
[]

`
11
subjects/binarycheck/README.md
Unescape
View File
@ -21,14 +21,15 @@ package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(BinaryCheck(5))
    fmt.Println(BinaryCheck(0))
    fmt.Println(BinaryCheck(8))
    fmt.Println(BinaryCheck(-9))
    fmt.Println(BinaryCheck(-4))
    fmt.Println(piscine.BinaryCheck(5))
    fmt.Println(piscine.BinaryCheck(0))
    fmt.Println(piscine.BinaryCheck(8))
    fmt.Println(piscine.BinaryCheck(-9))
    fmt.Println(piscine.BinaryCheck(-4))
}
```


14
subjects/binarycheck/main.go
Unescape
View File
@ -0,0 +1,14 @@
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BinaryCheck(5))
	fmt.Println(piscine.BinaryCheck(0))
	fmt.Println(piscine.BinaryCheck(8))
	fmt.Println(piscine.BinaryCheck(-9))
	fmt.Println(piscine.BinaryCheck(-4))
}
43
subjects/compare-exam/README.md
Unescape
View File
@ -1,43 +0,0 @@
## compare

## Instructions

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

import "fmt"

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
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
9
subjects/compare-exam/main.go
Unescape
View File
@ -1,9 +0,0 @@
package main

import "fmt"

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
7
subjects/compare/main.go
Unescape
View File
@ -2,10 +2,11 @@ package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "lut!"))
	fmt.Println(piscine.Compare("Ola!", "Ol"))
}

## concatslice

### Instructions

Write a function `ConcatSlice()` that takes two slices of integers as arguments and returns the concatenation of the two slices.

### Expected function

```go
func ConcatSlice(slice1, slice2 []int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.ConcatSlice([]int{1, 2, 3}, []int{}))
}
```

And its output:

```console
$ go run .
[1 2 3 4 5 6]
[4 5 6 7 8 9]
[1 2 3]
```
## digitlen

### Instructions

Write a function `DigitLen()` that takes two integers as arguments and returns the times the first `int` can be divided by the second until it reaches zero.

- The second `int` must be between **_2_** and **_36_**. If not, the function returns `-1`.
- If the first `int` is negative, reverse the sign and count the digits.

### Expected function

```go
func DigitLen(n, base int) int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
```

And its output:

```console
$ go run . | cat -e
3$
7$
2$
-1$
```
## hashcode

### Instructions

Write a function called `HashCode()` that takes a `string` as an argument and returns a new **hashed** `string`.

- The hash equation is computed as follows:

`(ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.`

- If the resulting character is unprintable add `33` to it.

### Expected function

```go
func HashCode(dec string) string {
}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
```

And its output:

```console
$ go run .
B
CD
EDF
Spwwz+bz}wo
```





## fifthandskip
### Instructions
- Write a function FifthAndSkip() that takes a string and returns another string. The function separates every five characters of the string with a space and removes the sixth one.

If there are spaces in the middle of a word, ignore them and get the first character after the spaces until you reach a length of 5.
If the string is less than 5 characters return Invalid Input followed by a newline \n.
If the string is empty return a newline \n.
#### Expected function
`
func FifthAndSkip(str string) string {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(piscine.FifthAndSkip("This is a short sentence"))
	fmt.Print(piscine.FifthAndSkip("1234"))
}
`
And its output:
`
$ go run . | cat -e
abcde ghijk mnopq stuwx z$
Thisi ashor sente ce$
Invalid Input$
`
7
subjects/fifthandskip/README.md
Unescape
View File
@ -25,12 +25,13 @@ package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(piscine.FifthAndSkip("This is a short sentence"))
	fmt.Print(piscine.FifthAndSkip("1234"))
}
```


12
subjects/fifthandskip/main.go
Unescape
View File
@ -0,0 +1,12 @@
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(piscine.FifthAndSkip("This is a short sentence"))
	fmt.Print(piscine.FifthAndSkip("1234"))
}
44
subjects/firstrune-exam/README.md
Unescape
View File
@ -1,44 +0,0 @@
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
)

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
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
14
subjects/firstrune/main.go
Unescape
View File
@ -0,0 +1,14 @@
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
7

## fishandchips
### Instructions
- Write a function called FishAndChips() that takes an int and returns a string.

If the number is divisible by 2, print fish.
If the number is divisible by 3, print chips.
If the number is divisible by 2 and 3, print fish and chips.
If the number is negative return error: number is negative.
If the number is non divisible by 2 or 3 return error: non divisible.

#### Expected function
`
func FishAndChips(n int) string {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FishAndChips(4))
	fmt.Println(piscine.FishAndChips(9))
	fmt.Println(piscine.FishAndChips(6))
}
`
And its output:
`
$ go run . | cat -e
fish$
chips$
fish and chips
`

## fromto
### Instructions
Write a function that takes two integers and returns a string showing the range of numbers from the first to the second.

The numbers must be separated by a comma and a space.
If any of the arguments is bigger than 99 or less than 0, the function returns Invalid followed by a newline \n.
Prepend a 0 to any number that is less than 10.
Add a new line \n at the end of the string.
#### Expected function
`
func FromTo(from int, to int) string {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
	"fmt"
	"psicine"
)

func main() {
	fmt.Print(piscine.FromTo(1, 10))
	fmt.Print(piscine.FromTo(10, 1))
	fmt.Print(piscine.FromTo(10, 10))
	fmt.Print(piscine.FromTo(100, 10))
}
`
And its output:
`
$ go run . | cat -e
01, 02, 03, 04, 05, 06, 07, 08, 09, 10$
10, 09, 08, 07, 06, 05, 04, 03, 02, 01$
10$
Invalid$
`

## issorted
### Instructions
W- rite a function IsSorted() that returns true, if the slice of int is sorted, otherwise returns false.

The function passed as an argument func(a, b int) returns a positive int if the first argument is greater than the second argument, it returns 0 if they are equal and it returns a negative int otherwise.

To do your testing you have to write your own f function.

### Expected function
`
func IsSorted(f func(a, b int) int, a []int) bool {

}
`
Usage
Here is a possible program to test your function (without f):
`
package main

import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
`
And its output:
`
$ go run .
true
false
$
`
### Notions
Function literals
Function declaration
Function types


44
subjects/lastrune-exam/README.md
Unescape
View File
@ -1,44 +0,0 @@
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
)

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
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

102
subjects/listremoveif-exam/README.md
Unescape
View File
@ -1,102 +0,0 @@
## listremoveif

### Instructions

Write a function `ListRemoveIf` that removes all elements that are equal to the `data_ref` in the argument of the function.

### Expected function and structure

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link, 1)
	ListPushBack(link, "Hello")
	ListPushBack(link, 1)
	ListPushBack(link, "There")
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, "How")
	ListPushBack(link, 1)
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

```

And its output :

```console
$ go run .
----normal state----
1 -> <nil>
------answer-----
<nil>

----normal state----
1 -> Hello -> 1 -> There -> 1 -> 1 -> How -> 1 -> are -> you -> 1 -> <nil>
------answer-----
Hello -> There -> How -> are -> you -> <nil>
$
```
55
subjects/listremoveif-exam/main.go
Unescape
View File
@ -1,55 +0,0 @@
package main

import "fmt"

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link, 1)
	ListPushBack(link, "Hello")
	ListPushBack(link, 1)
	ListPushBack(link, "There")
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, "How")
	ListPushBack(link, 1)
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
37
subjects/max-exam/README.md
Unescape
View File
@ -1,37 +0,0 @@
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

import "fmt"

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
```

And its output :

```console
$ go run .
123
$
```
12
subjects/max/main.go
Unescape
View File
@ -0,0 +1,12 @@
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
43
subjects/printcomb-exam/README.md
Unescape
View File
@ -1,43 +0,0 @@
## printcomb

### Instructions

Write a function that prints, in ascending order and on a single line: all **unique** combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.

These combinations are separated by a comma and a space.

### Expected function

```go
func PrintComb() {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
	PrintComb()
}
```

This is the incomplete output :

```console
$ go run . | cat -e
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789$
$
```

- `000` or `999` are not valid combinations because the digits are not different.

- `987` should not be shown because the first digit is not less than the second.

### Notions

- [01-edu/z01](https://github.com/01-edu/z01)
7
subjects/printcomb/main.go
Unescape
View File
@ -0,0 +1,7 @@
package main

import "piscine"

func main() {
	piscine.PrintComb()
}
4
subjects/printmemory/README.md
Unescape
View File
@ -23,8 +23,10 @@ Here is a possible program to test your function :
```go
package main

import "piscine"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	piscine.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
```


7
subjects/printmemory/main.go
Unescape
View File
@ -0,0 +1,7 @@
package main

import "piscine"

func main() {
	piscine.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
9

## retainfirsthalf
#### Instructions
- Write a function called RetainFirstHalf() that takes a string as an argument and returns the first half of this string.

If the length of the string is odd, round it down.
If the string is empty, return an empty string.
If the string length is equal to one, return the string.
#### Expected function
`
func RetainFirstHalf(str string) string {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(piscine.RetainFirstHalf("A"))
	fmt.Println(piscine.RetainFirstHalf(""))
	fmt.Println(piscine.RetainFirstHalf("Hello World"))
}
`
And its output:
`
$ go run . | cat -e
This is the 1st half$
A$
$
Hello$
`

9

## revconcatalternate
#### Instructions
- Write a function RevConcatAlternate() that receives two slices of int as arguments and returns a new slice with alternated values of each slice in reverse order.

The input slices can have different lengths.
The new slice should start with the elements from the largest slice first and when they became equal size slices, it should add an element of the first given slice.
If the slices are of equal length, the new slice should start with an element of the first slice.
Note: you can check the examples bellow for more details.

# Expected function
`
func RevConcatAlternate(slice1,slice2 []int) []int {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(piscine.RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
`
And its output:
`
$ go run .
[3 6 2 5 1 4]
[9 8 7 3 6 2 5 1 4]
[8 9 3 2 5 1 4]
[3 2 1]
`

45
subjects/rot14-exam/README.md
Unescape
View File
@ -1,45 +0,0 @@
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

```go
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
```

And its output :

```console
$ go run .
Vszzc! Vck ofs Mci?
$
```

## saveandmiss
### Instructions
- Write a function called SaveAndMiss() that takes a string and an int as an argument. The function should move through the string in sets determined by the int, saving the first set, omitting the second, saving the third, and so on, in a 'save' and 'miss' fashion until the end of the string is reached. Return a string containing the saved characters.

If the int is 0 or a negative number return the original string.

Expected function
`
func SaveAndMiss(arg string, num int) string {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SaveAndMiss("123456789", 3))
	fmt.Println(piscine.SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(piscine.SaveAndMiss("", 3))
	fmt.Println(piscine.SaveAndMiss("hello you all ! ", 0))
	fmt.Println(piscine.SaveAndMiss("what is your name?", 0))
	fmt.Println(piscine.SaveAndMiss("go Exercise Save and Miss", -5))
}
`
And its output:
`
$ go run . | cat -e
123789$
abcghimnostuz$
$
hello you all ! $
what is your name?$
go Exercise Save and Miss$
`

## setspace
#### Instructions
- Write a function that takes a PascalCase string and returns the same string with a space between each word.

The function must return an empty string if the input is an empty string.

The function must return "Error" if the input string is not a PascalCase string.

The PascalCase begins with a capital letter, and each word begins with a capital letter without a space between them, for example: "HelloWorld" is a valid PascalCase string.

The PascalCase cannot contain any non-alphabetic character, for example: "Hello World12" is not a valid PascalCase string.

#### Expected function
`
func SetSpace(s string) string {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
	"fmt"
)

func main() {
	fmt.Println(SetSpace("HelloWorld"))
	fmt.Println(SetSpace("HelloWorld12"))
	fmt.Println(SetSpace("Hello World"))
	fmt.Println(SetSpace(""))
	fmt.Println(SetSpace("LoremIpsumWord"))
}
`

And its output:
`
$ go run . | cat -e
Hello World$
Error$
Error$
$
Lorem Ipsum Word$
$
`

## slice

### Instructions

The function receives a slice of strings and one or more integers, and returns a slice of strings. The returned slice is part of the received one but cut from the position indicated in the first int, until the position indicated by the second int.

In case there only exists one int, the resulting slice begins in the position indicated by the int and ends at the end of the received slice.

The integers can be negative.

### Expected function

```go
func Slice(a []string, nbrs... int) []string{

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

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", piscine.Slice(a, 1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
    fmt.Printf("%#v\n", piscine.Slice(a, -3))
    fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))
}
```

```console
$ go run .
[]string{"algorithm", "ascii", "package", "golang"}
[]string{"ascii", "package"}
[]string{"ascii", "package", "golang"}
[]string{"package"}
[]string(nil)
```

## sliceadd
### Instructions
- Write a function that takes a slice of integers and an int as arguments, adds the int to the slice and returns it.

If the slice is empty, return a slice with the new value.
#### Expected function
`
func SliceAdd(slice []int , num int) []int {

}
`
Usage
Here is a possible program to test your function:
`
package main

import (
	"fmt"
)

func main() {
	fmt.Println(SliceAdd([]int{1, 2, 3}, 4))
	fmt.Println(SliceAdd([]int{}, 4))
}
`
### And its output:
`
$ go run .
[1 2 3 4]
[4]
`

## sliceremove

### Instructions

Write a function that takes a slice of integers and an `int` as arguments, if the given `int` exists in the slice then remove it, otherwise return the slice.

- If the slice is empty, return the slice itself.

### Expected function

```go
func SliceRemove(slice []int , num int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(SliceRemove([]int{1, 2, 3}, 2))
	fmt.Println(SliceRemove([]int{4, 3}, 4))
	fmt.Println(SliceRemove([]int{}, 1))

}
```

And its output:

```console
$ go run .
[1 3]
[3]
[]
```
38
subjects/sortwordarr-exam/README.md
Unescape
View File
@ -1,38 +0,0 @@
## sortwordarr

### Instructions

Write a function `SortWordArr` that sorts by `ascii` (in ascending order) a `string` slice.

### Expected function

```go
func SortWordArr(a []string) {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}
```

And its output :

```console
$ go run .
[1 2 3 A B C a b c]
$
```
9
subjects/sortwordarr-exam/main.go
Unescape
View File
@ -1,9 +0,0 @@
package main

import "fmt"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	fmt.Println(result)
}
36
subjects/split-exam/README.md
Unescape
View File
@ -1,36 +0,0 @@
## split

### Instructions

Write a function that receives a string and a separator and returns a `slice of strings` that results of splitting the string `s` by the separator `sep`.

### Expected function

```go
func Split(s, sep string) []string {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
```

And its output :

```console
$ go run .
[]string{"Hello", "how", "are", "you?"}
$
```
11
subjects/split/main.go
Unescape
View File
@ -0,0 +1,11 @@
package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
}
36
subjects/strlen-exam/README.md
Unescape
View File
@ -1,36 +0,0 @@
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

import "fmt"

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
```

And its output :

```console
$ go run .
12
$
```
11
subjects/strlen/main.go
Unescape
View File
@ -0,0 +1,11 @@
package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
9
subjects/wordflip/README.md
Unescape
View File
@ -25,13 +25,14 @@ package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
	fmt.Print(piscine.WordFlip("First second last"))
	fmt.Print(piscine.WordFlip(""))
	fmt.Print(piscine.WordFlip("     "))
	fmt.Print(piscine.WordFlip(" hello  all  of  you! "))
}
```


13
subjects/wordflip/main.go
Unescape
View File
@ -0,0 +1,13 @@
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.WordFlip("First second last"))
	fmt.Print(piscine.WordFlip(""))
	fmt.Print(piscine.WordFlip("     "))
	fmt.Print(piscine.WordFlip(" hello  all  of  you! "))
}

## sliceadd

### Instructions

Write a function that takes a slice of integers and an `int` as arguments, adds the `int` to the slice and returns it.

- If the slice is empty, return a slice with the new value.

### Expected function

```go
func SliceAdd(slice []int , num int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(SliceAdd([]int{1, 2, 3}, 4))
	fmt.Println(SliceAdd([]int{}, 4))
}
```

And its output:

```console
$ go run .
[1 2 3 4]
[4]
```

## secondhalf

### Instructions

Write a function `SecondHalf()` that receives a slice of `int` and returns another slice of `int` with the second half of the values.

> If the length is odd, include the middle value in the result.

### Expected function

```go
func SecondHalf(slice []int) []int {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.SecondHalf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(piscine.SecondHalf([]int{1, 2, 3}))
}
```

And its output:

```console
$ go run . | cat -e
[6 7 8 9 10]
[2 3]
```

