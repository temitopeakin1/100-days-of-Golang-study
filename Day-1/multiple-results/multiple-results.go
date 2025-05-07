package main

import "fmt"

// what the function does
func swap(x, y string) (string, string) {
	return y, x
}

// the end of the function
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
