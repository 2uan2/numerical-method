package interpolation

import (
	"fmt"
)

type Point2D [2]float64

type Point3D [3]float64

func LinearLagrange(x float64, a float64, b float64, f func(float64) float64) float64 {
	f0 := f(a)
	fmt.Printf("f(a) = %f\n", f0)
	f1 := f(b)
	fmt.Printf("f(b) = %f\n", f1)
	L0 := (x - b) / (a - b)
	fmt.Printf("L(a) = %f\n", L0)
	L1 := (x - a) / (b - a)
	fmt.Printf("L(b) = %f\n", L1)
	return L0*f0 + L1*f1
}

func GeneralLagrange2DContinuous(xs []float64, ys []float64, stepAmount float64) []Point2D {
	minX := xs[0]
	maxX := xs[len(xs)-1]
	for _, x := range xs {
		if x < minX {
			minX = x
		} else if x > maxX {
			maxX = x
		}
	}

	points := make([]Point2D, 0)

	for x := minX; x < maxX; x += stepAmount {
		y := GeneralLagrange2D(x, xs, ys)
		points = append(points, Point2D{x, y})
	}

	return points
}

func GeneralLagrange2D(x float64, xs []float64, ys []float64) float64 {
	p := 0.0
	for idx, _ := range xs {
		// p += (l(idx, x, xs) / l(idx, curX, xs)) * ys[idx]
		p += l(idx, x, xs) * ys[idx]
		// fmt.Printf("L%d = %f\n", idx, l(idx, x, xs)/l(idx, curX, xs))
	}
	return p
}

func GeneralLagrange3D(x float64, xs []float64, ys []float64) float64 {
	p := 0.0
	for idx, curX := range xs {
		p += (l(idx, x, xs) / l(idx, curX, xs)) * ys[idx]
		// p += l(idx, x, xs) * ys[idx]
		// fmt.Printf("L%d = %f\n", idx, l(idx, x, xs)/l(idx, curX, xs))
	}
	return p
}

func l(i int, x float64, others []float64) float64 {
	l := 1.0
	for idx, curX := range others {
		if idx == i {
			continue
		}

		// l *= (x - curX)
		l *= (x - curX) / (others[i] - curX)
		// fmt.Printf("l *= (%f - %f)\n", x, curX)
	}
	// fmt.Printf("l%d(%f) = %f\n", i, x, l)
	return l
}
