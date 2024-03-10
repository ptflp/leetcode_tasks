package main

import "math"

func climbStairs(n int) int {
	n++
	sqrt5 := math.Sqrt(5)
	phi := (1 + sqrt5) / 2
	fib := (math.Pow(phi, float64(n)) - math.Pow(1-phi, float64(n))) / sqrt5

	return int(math.Round(fib))
}
