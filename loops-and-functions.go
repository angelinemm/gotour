package main

import (
	"fmt"
	"math"
	"strconv"
)

func MySqrt(x float64) float64 {
    z := float64(1)
	prev_z := z
	for {
	    z -= (z*z - x) / (2*z)
		if math.Abs(z - prev_z) < 0.00000000000001 {
		    return z
		} else {
		    prev_z = z
		}
	}
	return z
}

func main() {
    num := float64(2)
	fmt.Println(MySqrt(num))
	fmt.Println(math.Sqrt(num))
}