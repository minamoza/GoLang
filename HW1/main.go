package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	cnt := 0
	for {
		cnt++
		prevZ := z
		z -= (z*z - x) / (2*z)
		if z * z <= x || prevZ == z{
			break
		}
	}
	return z, cnt
}

func main() {
	ans, itr := Sqrt(3)
	fmt.Println("Answer from math.Sqrt: ", math.Sqrt(3))
	fmt.Printf("My answer is: %v \nNumber of iterations: %v\n", ans, itr)
}
