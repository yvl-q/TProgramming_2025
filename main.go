package main

import (
	"fmt"
	"math"
)

func main() {
	xStart := 0.2
	xEnd := 2.2
	dx := 0.4

	for x := xStart; x <= xEnd+0.0001; x += dx {
		var y float64

		if math.Abs(x) >= 1 {
			y = math.Pow(1.2, x) - math.Sqrt(x)
		} else {
			y = math.Acos(x)
		}

		fmt.Printf("x = %.1f, y = %.4f\n", x, y)
	}
}
