package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 4{
		return
	}
	str := os.Args[1]
	c1 := os.Args[2][0]
	c2 := os.Args[3][0]
	res := ""
	for _,c := range str{
		if c == rune(c1) {
			res += string(c2)
		}else{
			res += string(c)
		}
	}
	fmt.Println(res)
}

/*
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

*/