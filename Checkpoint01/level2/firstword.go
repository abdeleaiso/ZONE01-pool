package piscine

func FirstWord(s string) string {
	if len(s) == 0{
		return "\n"
	}
	first, last := 0, 0
	for i:=0; i<len(s); i++{
		if s[i] == ' '{
			continue
		}else {
			first = i
			break
		}
	}
	for i:=first; i<len(s);i++{
		if s[i] == ' '{
			last = i
			break
		}else{
			last++
		}
	}
	return s[first:last]+"\n"
}
/*
## firstword

### Instructions

Write a function that takes a string and return a string containing its first word, followed by a newline (`'\n'`).

- A word is a sequence of characters delimited by spaces or by the start/end of the argument.

### Expected Function

```go
func FirstWord(s string) string {
    // ...
}
```

### Usage

Here is a possible way to test your function:

```go
package main

import (
    "fmt"

    "piscine"
)

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}
```

And its output:

```console
$ go run .
hello

hello
$
```

*/