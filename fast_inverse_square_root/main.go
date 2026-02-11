package main

import (
	"fmt"
	"math"
)

// QUAKE 3: THE LEGEND. HACK THE HARDWARE CODE to optimize lights and shadows
func FastInverseSqrt(x float32) float32 {
	x2 := x * 0.5
	y := x

    // The "Evil Bit Hack" in Go 
    i := math.Float32bits(y) // Get bits as uint32
    i = 0x5f3759df - (i >> 1) // The Magic Number shift 
    y = math.Float32frombits(i) // Convert bits back to float32

    // One iteration of Newton's Method 
    y = y * (1.5 - (x2 * y * y))

    return y
}

func main() {
    n := float32(2.0) 
    fmt.Printf("Fast Inverse Sqrt of 2: %v\n", FastInverseSqrt(n))
    fmt.Printf("Math.Sqrt version (1/sqart(2)): %v\n", 1/math.Sqrt(float64(n)))
}
