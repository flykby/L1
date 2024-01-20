package main

import "fmt"



func main() {
	a := 4
	b := 1
	fmt.Printf("a: %d b: %d\n", a, b)

	a, b = b, a
	
	fmt.Printf("a: %d b: %d\n", a, b)
	
	a ^= b
	b ^= a
	a ^= b
	
	fmt.Printf("a: %d b: %d\n", a, b)
}