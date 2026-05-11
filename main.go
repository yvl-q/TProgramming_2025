package main

import (
	"fmt"
	"math"
)

func calcY(x float64) float64 {
	if math.Abs(x) >= 1 {
		return math.Pow(1.2, x) - math.Sqrt(x)
	}

	return math.Acos(x)
}

func main() {
	fmt.Println("Задача A")

	xStart := 0.2
	xEnd := 2.2
	dx := 0.4

	for x := xStart; x <= xEnd+0.0001; x += dx {
		y := calcY(x)
		fmt.Printf("x = %.1f, y = %.4f\n", x, y)
	}

	fmt.Println()
	fmt.Println("Задача B")

	xValues := []float64{0.1, 0.9, 1.2, 1.5, 2.3}

	for i := 0; i < len(xValues); i++ {
		x := xValues[i]
		y := calcY(x)
		fmt.Printf("x = %.1f, y = %.4f\n", x, y)
	}
}
