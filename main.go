package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0.2
	var y float64

	if math.Abs(x) >= 1 {
		y = math.Pow(1.2, x) - math.Sqrt(x)
	} else {
		y = math.Acos(x)
	}

	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
