/*
## weareunique

### Instructions

Write a function that takes two `strings`'s and returns the number of characters that are not included in both, without repeating characters.

- If there is no unique characters return `0`.
- If both strings are empty return `-1`.

### Expected function

```go
func WeAreUnique(str1 , str2 string) int {

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
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
```

And its output:

```console
$ go run .
2
-1
6
```
*/

package piscine

//merthod1 usnig slice
func WeAreUnique1(str1 , str2 string) int {
	if str1 == "" && str2 == ""{
		return -1
	}
	total := str1 + str2
	checker := [256]int{0}
	count := 0
	for i:=0;i<len(total);i++{
		checker[total[i]]++
	}
	for i:=0; i<len(checker);i++{
		if checker[i] == 1{
			count++
		}
	}
	return count
}

//method2 using map
func WeAreUnique2(str1 , str2 string) int {
	if str1 == "" && str2 == ""{
		return -1
	}
	total := str1 + str2
	m := make(map[string]int)
	count := 0
	for i:=0;i<len(total);i++{
		m[string(total[i])]++
	}
	for _,v := range m{
		if v == 1{
			count++
		}
	}
	return count
}