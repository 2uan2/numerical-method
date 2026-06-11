package main

import (
	// "algorithms/newton"
	// "algorithms/interpolation"

	"algorithms/ode"
	"fmt"
	"math"
)

var (
	TOL = 0.00001
	n0  = 100
)

func f(x float64) float64 {
	// return x*x - x - 1.0
	// return x*x*x - 2*x + 1
	// return math.Cos(x)
	// return math.Sin(x)
	// return x*math.Pow(math.E, 1-(x/3)) - 0.75
	return x*x - 3*x + 1
}

func c(t float64) float64 {
	A := 3.0
	return A * t * math.Pow(math.E, -t/3.0)
}

func main() {
	// A0 := [][]int{
	// 	{2, 1, 1, 0},
	// 	{4, 3, 3, 1},
	// 	{8, 7, 9, 5},
	// 	{6, 7, 9, 8},
	// }
	// b := []int{6, 15, 41, 40}
	// A1 := [][]float64{
	// 	{1, 0, -1},
	// 	{-0.5, 1, -0.25},
	// 	{1, -0.5, 1},
	// }
	// b1 := []float64{0.2, -1.425, 2}
	// x := []float64{0, 0, 0}
	// threshold := 0.01
	// A2 := [][]float64{
	// 	{1, -0.25, -0.25, 0},
	// 	{-0.25, 1, 0, -0.25},
	// 	{-0.25, 0, 1, -0.25},
	// 	{0, -0.25, -0.25, 1},
	// }
	// b2 := []float64{50, 50, 25, 25}
	// x := []float64{100, 100, 100, 100}
	// fmt.Println(gauss.GaussSeidel(A1, b1, x, 300, threshold))
	// fmt.Println(math.Log(9.0))
	// f := func(x float64) float64 {
	// 	return math.Log(x)
	// }
	// fmt.Println(interpolation.LinearLagrange(9.2, 9.0, 9.5, f))
	// xs := []float64{1950, 1960, 1970, 1980, 1990, 2000}
	// ys := []float64{151326, 179323, 203302, 226542, 249633, 281422}
	// xs := []float64{0.0, 0.2, 0.4, 0.6, 0.8, 1.0}
	// ys := []float64{0.0, 0.2227, 0.4284, 0.6039, 0.7422, 0.8426}
	// fmt.Println(interpolation.GeneralLagrange(0.33, xs, ys))
	var x0 float64 = 0
	var y0 float64 = 0.5
	h := 0.2
	f := func(x float64, y float64) float64 {
		// return 1.0 + math.Pow(x, 2) + y
		// return 1 + math.Pow(y, 2) + math.Pow(x, 3)
		return 1 + y + math.Pow(x, 2)
	}

	// fmt.Println(ode.Heun(f, x0, y0, h, 3))
	fmt.Println(ode.RungeKutta4(f, x0, y0, h, 3))
	// cur_max := 0.0
	// max_day := 0
	// for i := 0.0; i < 28.0; i++ {
	// 	if interpolation.GeneralLagrange(i, xs, ys) > cur_max {
	// 		cur_max = interpolation.GeneralLagrange(i, xs, ys)
	// 		max_day = int(i)
	// 	}
	// }
	// fmt.Println(cur_max)
	// fmt.Println(max_day)
}

// f0 ⫽ f (0) ⫽ 1, f1 ⫽ f (2) ⫽ 9, f2 ⫽ f (4) ⫽ 41,
// f3 ⫽ f (6) ⫽ 41, k 0 ⫽ 0, k 3 ⫽ ⫺12
