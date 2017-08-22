package main

import (
	"fmt"
	"math"
	"strconv"
)

// MySqrt approximate the value of sqrt of x
func MySqrt(x float64) float64 {
    z := float64(1)
	prevZ := z
	for {
	    z -= (z*z - x) / (2*z)
		if math.Abs(z - prevZ) < 0.00000000000001 {
		    return z
		}
		prevZ = z
	}
	return z
}

func main() {
    num := float64(2)
	fmt.Println(MySqrt(num))
	fmt.Println(math.Sqrt(num))
}
