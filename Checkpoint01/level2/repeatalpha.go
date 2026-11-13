package piscine

func RepeatAlpha(s string) string {
	l := 0
	str := ""
	for _, c := range s{
		if c >= 'a' && c <= 'z'{
			l = int(c-'a')+1
		}else if c >= 'A' && c <= 'Z'{
			l = int(c-'A')+1
		}else { 
			str += string(c)
		}
		for i:=0; i<l; i++{
			str += string(c)
		}
		l = 0
	}
	return str
}

/*
## repeatalpha

### Instructions

Write a function called `RepeatAlpha` that takes a `string` and displays it repeating each alphabetical character as many times as its alphabetical index.

`'a'` becomes `'a'`, `'b'` becomes `'bb'`, `'e'` becomes `'eeeee'`, etc...

### Expected Function

```go
func RepeatAlpha(s string) string {
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
	fmt.Println(epeatAlpha("abc"))
	fmt.Println(epeatAlpha("Choumi."))
	fmt.Println(epeatAlpha(""))
	fmt.Println(epeatAlpha("abacadaba 01!"))
}
```

And its output:

```console
$ go run . | cat -e
abbccc$
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
$
abbacccaddddabba 01!$
$
```

*/