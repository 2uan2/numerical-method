package main

import (
	"algorithms/interpolation"
	"fmt"
	"math"
	"slices"
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
	// gause seidel and stuff
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

	// runge kutta and other first order ODE
	// var x0 float64 = 0
	// var y0 float64 = 2.0
	// h := 0.1
	// f := func(x float64, y float64) float64 {
	// 	// return 1.0 + math.Pow(x, 2) + y
	// 	// return 1 + math.Pow(y, 2) + math.Pow(x, 3)
	// 	// return 1 + y + math.Pow(x, 2)
	// 	return 2*x + math.Pow(y, 2)*x
	// }
	// // fmt.Println(ode.Heun(f, x0, y0, h, 3))
	// x1, y1 := ode.Heun(f, x0, y0, h, 1)
	// fmt.Println("   	 2.5831917308068606")
	// for i := range 10 {
	// 	_, y := ode.HeunMultistep(f, x0, y0, x1, y1, h, 3, i)
	// 	fmt.Printf("k = %d -> %.16f\n", i, y) // 2.5831917308068606
	// }
	// fmt.Println(ode.RungeKutta4(f, x0, y0, h, 3))

	// ODE system
	// fs := make([]func(args ...float64) float64, 2)
	// fs[0] = func(args ...float64) float64 {
	// 	return args[2]
	// 	// t := args[0]
	// 	// return 3*args[1] + 2*args[2] - (math.Pow(t, 2)*2+1)*math.Pow(math.E, 2*t)
	// }
	// fs[1] = func(args ...float64) float64 {
	// 	return 4 * (args[1] - args[0])
	// 	// t := args[0]
	// 	// return 4*args[1] + args[2] + (math.Pow(t, 2)+2*t-4)*math.Pow(math.E, 2*t)
	// }
	// x: args[0], y: args[1], y': args[2]
	// fs[2] = func(args ...float64) float64 {
	// 	return -5*args[2] - 2*args[1] - 8*args[3]
	// }
	// fs[3] = func(args ...float64) float64 {
	// 	return args[4]
	// }
	// fs[4] = func(args ...float64) float64 {
	// 	return 2 - args[1] - 2*args[0]*args[3]
	// }
	// x0 := 0.0
	// first x0 is just x
	// x0 := 0.0
	// x1 := 1.0
	// vars := []float64{0.0, 0.0, 2.0000}
	// h := 0.01

	// fmt.Println(ode.RungeKuttaSystem(fs, x0, vars, h, 3))
	// fmt.Println(ode.Shooting(fs, x0, x1, vars, h, 2))
	// fmt.Println(interpolation.GeneralLagrange2DContinuous([]float64{-0.7688, 0.99}, []float64{0, 1}, 0.1))

	// interpolation with lagrange type shi
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
	points := []interpolation.Point2D{{1.0, 2.0}, {2.0, 4.5}, {3.0, 5.0}, {5.0, 3.0}, {3.0, 2.0}}
	// f := interpolation.Spline(points)
	// fmt.Println(f(t))
	xs := slices.Collect(func(yield func(float64) bool) {
		for _, point := range points {
			if !yield(point[0]) {
				return
			}
		}
	})
	ys := slices.Collect(func(yield func(float64) bool) {
		for _, point := range points {
			if !yield(point[1]) {
				return
			}
		}
	})
	fX := interpolation.CatmullRomSpline(xs)
	fmt.Println("done")
	fY := interpolation.CatmullRomSpline(ys)
	// t := 1.0
	for t := 0.0; t < 4.0; t += 0.1 {
		fmt.Println(t, fX(t), fY(t))
	}

	// A0 := [][]float64{
	// 	{-3, 2, -1},
	// 	{6, -6, 7},
	// 	{3, -4, 4},
	// }
	// b := []float64{-1, -7, -6}
	// fmt.Println(gauss.GaussElimination(A0, b))

}

// f0 ⫽ f (0) ⫽ 1, f1 ⫽ f (2) ⫽ 9, f2 ⫽ f (4) ⫽ 41,
// f3 ⫽ f (6) ⫽ 41, k 0 ⫽ 0, k 3 ⫽ ⫺12
