package ode

import (
	"algorithms/interpolation"
	"fmt"
	"math"
)

func Midpoint(f func(a, b float64) float64, x0, y0, h float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var xMid float64
	var yMid float64

	for range step {
		yMid = yCur + (h/2)*f(xCur, yCur)
		xMid = xCur + h/2
		yCur = yCur + h*f(xMid, yMid)
		xCur = xCur + h
		fmt.Printf("y(%f) = %f\n", xCur, yCur)
	}

	return xCur, yCur
}

func MidPointSystem(fs []func(args ...float64) float64, x float64, initials []float64, h float64, step int) []float64 {
	return []float64{}
}

func RungeKuttaSystem(fs []func(args ...float64) float64, x float64, initials []float64, h float64, step int) []float64 {
	xCur := x
	zCur := initials
	zNew := make([]float64, len(initials))
	xAmount := len(initials) - len(fs)
	for range step {
		for i, f := range fs {
			k1 := f(zCur...)
			k2Args := make([]float64, len(zCur))
			k2Args[0] = zCur[0] + h
			for i := 1; i < len(zCur); i++ {
				k2Args[i] = zCur[i] + h*k1
			}
			k2 := f(k2Args...)

			zNew[i+xAmount] = zCur[i+xAmount] + h/2.0*(k1+k2)
		}

		for i := range initials {
			zCur[i] = zNew[i]
		}
		xCur = xCur + h
		if xAmount >= 1 {
			zCur[0] = xCur
		}
		fmt.Printf("Z[%f] = %v\n", xCur, zCur)
	}

	return zCur
}

func EulerSystem(fs []func(args ...float64) float64, x float64, initials []float64, h float64, step int) []float64 {
	xCur := x
	zCur := initials
	zNew := make([]float64, len(initials))
	xAmount := len(initials) - len(fs)
	for range step {
		for i, f := range fs {
			zNew[i+xAmount] = zCur[i+xAmount] + h*f(zCur...)
		}

		for i := range initials {
			zCur[i] = zNew[i]
		}
		xCur = xCur + h
		if xAmount >= 1 {
			zCur[0] = xCur
		}
		fmt.Printf("Z[%f] = %v\n", xCur, zCur)
	}

	return zCur
}

func Shooting(fs []func(args ...float64) float64, xStart, xEnd float64, ys []float64, h float64, step int) []float64 {
	threshold := 0.01
	yActual := ys[len(ys)-1]
	// ys[1] = ys[len(ys)-2]
	ys[len(ys)-1] = ys[len(ys)-2]
	xGuesses := make([]float64, 0)
	yGuesses := make([]float64, 0)

	for i := range step {
		stepAmount := (xEnd - xStart) / h
		guessBoundary := RungeKuttaSystem(fs, xStart, ys, h, int(stepAmount))
		fmt.Println(guessBoundary)
		if math.Abs(guessBoundary[1]-yActual) < threshold {
			break
		}
		// first step
		if i == 0 {
			xGuesses = append(xGuesses, float64(i))
			yGuesses = append(yGuesses, guessBoundary[1])
			ys[len(ys)-1] = ys[len(ys)-1] + 1
		} else if i == 1 {
			xGuesses = append(xGuesses, float64(i))
			yGuesses = append(yGuesses, guessBoundary[1])
			ys[len(ys)-1] = interpolation.GeneralLagrange2D(yActual, yGuesses, xGuesses)
			fmt.Println("ys: ", ys)
		} else {

		}

	}

	return []float64{}
}
