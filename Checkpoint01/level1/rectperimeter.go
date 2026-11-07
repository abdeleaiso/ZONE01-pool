package piscine

func RectPerimeter(w, h int) int {
	if h <= 0 || w <= 0{
		return -1
	} 
	return 2*(h+w)
}

/*
## Instructions
Write a function that takes two int's as arguments, representing the length of width and height of a rectangle and returning the perimeter of the rectangle.

- If one of the arguments is negative it should return -1.

### Expected function:

func RectPerimeter(w, h int) int {

}

### Usage

package main

import (
	"fmt"
)

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}

- output:

$ go run .
24$
2666664$
-1$
$
*/
