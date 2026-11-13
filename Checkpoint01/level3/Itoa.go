package piscine

//methode 1
func Itoa(n int) string {
	if n == 0{
		return "0"
	}
	negative := false
	if n == -9223372036854775808{
		return "-9223372036854775808"
	}
	if n < 0 {
		negative = true
		n *= (-1)
	}
	arr := []int{}
	for n != 0{
		arr = append(arr, n%10)
		n /= 10
	}
	res := ""
	if negative{
		res += "-"
	}
	for i:=len(arr)-1; i>=0; i--{
		res += string(arr[i]+'0')
	}
	return res
}

//methode 2: reverse of method1
func Itoa2(n int) string {
	if n == 0 {
		return "0"
	}

	res := ""
	if n > 0 {
		n = -n
	} else {
		res += "-"
	}

	for n < 0 {
		digit := -(n % 10)
		res = res + string(rune(digit+'0'))
		n /= 10
	}

	if res[0] != '-' {
		runes := []rune(res)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		res = string(runes)
	}

	return res
}

/*
## itoa

### Instructions

- Write a function that simulates the behavior of the `Itoa` function in Go. `Itoa` transforms a number represented as an`int` in a number represented as a `string`.

- For this exercise the handling of the signs + or - **does have** to be taken into account.

## Expected function

```go
func Itoa(n int) string {

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
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}
```

And its output :

```console
$ go run .
12345
0
-1234
987654321
$
```

### Notions

- [strconv/Itoa](https://pkg.go.dev/strconv#Itoa)

*/