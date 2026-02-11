package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	delta := 0.000000000000000000001 // epsilon

	for i := range 10 {
		prevZ := z
		z -= (z*z - x) / (2 * z) // Newton's Method formula

		fmt.Printf("Iteration %d: %v\n", i+1, z)

		// Optimization: Stop if the change is smaller than delta
		if math.Abs(z-prevZ) < delta {
			fmt.Println("Converged early!")
			break
		}
	}
    return z
}

func main() {
    val := 2.0  
    fmt.Printf("Newton's Method result: %v\n", Sqrt(val))
    fmt.Printf("Standard Library result: %v\n", math.Sqrt(val))
}
