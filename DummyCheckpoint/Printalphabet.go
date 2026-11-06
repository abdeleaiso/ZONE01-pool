package main

import "github.com/01-edu/z01"

func main(){
	for i:='a'; i<='z';i++{
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

/*
printalphabet
Instructions
Write a program that prints the Latin alphabet in lowercase on a single line.

A line is a sequence of characters preceding the end of line character ('\n').

Usage
$ go run .
abcdefghijklmnopqrstuvwxyz
$
Notions
01-edu/z01
*/