package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return // known as "naked" return, because the return values are named
	// and the return statement does not specify them
}
func main() {
	fmt.Println(split(17))

}
