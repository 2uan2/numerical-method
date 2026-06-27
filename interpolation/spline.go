package interpolation

import (
	"algorithms/gauss"
	"fmt"
	"math"
	"slices"
)

func Spline(points []Point2D) func(t float64) float64 {
	slopes := make([]float64, 0)
	coefficients := make([][]float64, 0)

	var x, y, slope float64
	A0 := make([][]float64, 0)
	b := make([]float64, 0)
	for i := 0; i < len(points); i++ {
		// if first or if last
		x = points[i][0]
		y = points[i][1]

		if i == 0 {
			deltaF := math.Abs(points[i+1][1] - points[i][1]) // rise
			deltat := math.Abs(points[i+1][0] - points[i][0]) // run
			slope = deltaF / deltat
			// fmt.Println("slope at 0 ", slope)
		} else if i == len(points)-1 {
			deltaF := math.Abs(points[i][1] - points[i-1][1]) // rise
			deltat := math.Abs(points[i][0] - points[i-1][0]) // run
			slope = deltaF / deltat
			// fmt.Println("slope at last ", slope)
		} else {
			deltaF := math.Abs(points[i+1][1] - points[i-1][1]) // rise
			deltat := math.Abs(points[i+1][0] - points[i-1][0]) // run
			slope = deltaF / deltat
			// fmt.Printf("slope at %d: %v\n", i, slope)
		}
		slopes = append(slopes, slope)

		// add ax^3 + bx^2 + cx + d = y (y is the b append)
		A0 = append(A0, []float64{math.Pow(x, 3), math.Pow(x, 2), x, 1})
		b = append(b, y)

		// add 3x^2 + 2bx + c = y'
		A0 = append(A0, []float64{3 * math.Pow(x, 2), 2 * x, 1, 0})
		b = append(b, slope)

		fmt.Printf("---------i = %d--------\n", i)
		fmt.Println("now A0 is ", A0)
		fmt.Println("now b is ", b)
		// don't calculate coefficient for the first loop cause interval needs both ends calculated first
		if i == 0 {
			fmt.Println("continuing")
			continue
		}
		coef := gauss.GaussElimination(A0, b)
		fmt.Println("coef are ", coef)
		coefficients = append(coefficients, coef)

		// move matrix[2] and [3] down to [0] and [1]
		A0[0] = A0[2]
		A0[1] = A0[3]
		b[0] = b[2]
		b[1] = b[3]
		A0 = slices.Delete(A0, 1, 3)
		b = slices.Delete(b, 1, 3)
	}
	return func(t float64) float64 {
		// fmt.Println(slopes)
		for i := 0; i < len(points); i++ {
			if t >= points[i][0] && t <= points[i+1][0] {
				fmt.Printf("in range %f %f\n", points[i][0], points[i+1][0])
				a := coefficients[i][0]
				b := coefficients[i][1]
				c := coefficients[i][2]
				d := coefficients[i][3]
				fmt.Println("a = ", a)
				fmt.Println("b = ", b)
				fmt.Println("c = ", c)
				fmt.Println("d = ", d)
				return a*math.Pow(t, 3) + b*math.Pow(t, 2) + c*t + d
			}
		}
		return 1.0
	}
}

func CatmullRomSpline(points []float64) func(t float64) float64 {
	slopes := make([]float64, 0)
	coefficients := make([][]float64, 0)

	var x, slope float64
	A0 := make([][]float64, 0)
	b := make([]float64, 0)
	for t := 0; t < len(points); t++ {
		x = points[t]

		// if first or if last
		if t == 0 {
			slope = math.Abs(points[t+1] - x)
		} else if t == len(points)-1 {
			slope = math.Abs(x - points[t-1])
		} else {
			slope = math.Abs(points[t+1]-points[t-1]) / 2.0
		}
		slopes = append(slopes, slope)

		// add at^3 + bt^2 + ct + d = x (x is the b append, t = 0)
		A0 = append(A0, []float64{math.Pow(float64(t), 3), math.Pow(float64(t), 2), float64(t), 1})
		b = append(b, x)

		// add 3t^2 + 2bt + c = x'
		A0 = append(A0, []float64{3 * math.Pow(float64(t), 2), 2 * float64(t), 1, 0})
		b = append(b, slope)

		fmt.Printf("---------i = %d--------\n", t)
		fmt.Println("now A0 is ", A0)
		fmt.Println("now b is ", b)
		// don't calculate coefficient for the first loop cause interval needs both ends calculated first
		if t == 0 {
			// fmt.Println("continuing")
			continue
		}
		coef := gauss.GaussElimination(A0, b)
		fmt.Printf("coef are %.2f*t^3 + %.2f*t^2 + %.2f*t + %.2f\n", coef[0], coef[1], coef[2], coef[3])
		coefficients = append(coefficients, coef)

		// move matrix[2] and [3] down to [0] and [1]
		A0[0] = A0[2]
		A0[1] = A0[3]
		b[0] = b[2]
		b[1] = b[3]
		A0 = slices.Delete(A0, 1, 3)
		b = slices.Delete(b, 1, 3)
	}
	return func(t float64) float64 {
		for i := 0; i < len(points)-1; i++ {
			if t >= float64(i) && t <= float64(i+1) {
				// fmt.Printf("in range %f %f\n", points[i], points[i+1])
				a := coefficients[i][0]
				b := coefficients[i][1]
				c := coefficients[i][2]
				d := coefficients[i][3]
				// fmt.Println("a = ", a)
				// fmt.Println("b = ", b)
				// fmt.Println("c = ", c)
				// fmt.Println("d = ", d)
				return a*math.Pow(t, 3) + b*math.Pow(t, 2) + c*t + d
			}
		}
		return 1.0
	}
}

// a + b + c + d = 2
// 3a + 2b + c = 2.5
// 8a + 4b + 2c + d = 4.5
// 12a + 4b + c = 1.5
