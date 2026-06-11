package main

import (
	"fmt"
	"math"
)

func bisection(start float64, end float64, f func(float64) float64) float64 {
	mid := (start + end) / 2
	f1 := f(start)
	f2 := f(mid)
	f3 := f(end)
	fmt.Println("-------")
	fmt.Println("f1: ", f1)
	fmt.Println("f2: ", f2)
	fmt.Println("f3: ", f3)

	if math.Abs(f2) < 0.01 {
		return mid
	}

	if f1*f2 > 0 {
		return bisection(mid, end, f)
	} else {
		return bisection(start, mid, f)
	}
}

func backwardI(x float64) float64 {
	var result float64
	var ans float64
	if x == 15 {
		result = 0.0
		ans = 0.0
	} else {
		result = backwardI(x + 1)
		ans = (math.E - result) / (x + 1)
	}
	fmt.Printf("(%f - %f) / (%f + 1) = %f\n", math.E, result, x, ans)
	return ans
}

func forwardI(x float64) float64 {
	decimal := 5
	if x == 0 {
		return round(math.E, decimal) - 1.0
	}
	return round(math.E, decimal) - round(x*forwardI(x-1), decimal)
}

func round(val float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(val*shift) / shift
}
