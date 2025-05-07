package main

import "fmt"

// initialise global variables
var i, j int = 1, 2

func main() {
	// initialise local variables
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
