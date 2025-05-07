// The expression T(v) converts the value v to the type T.

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // take note of what happens here
	// x*x + y*y is an int, but math.Sqrt expects a float64
	// So we need to convert x*x + y*y to a float64
	// by using the conversion function float64(x*x + y*y)
	// The conversion function takes a single argument and returns a value of the specified type.
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
