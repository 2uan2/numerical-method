package ode

import "fmt"

func RungeKutta2Alpha(f func(a float64, b float64) float64, x0 float64, y0 float64, h float64, alpha float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var k1 float64
	var k2 float64
	for range step {
		k1 = h * f(xCur, yCur)
		k2 = h * f(xCur+alpha*h, yCur+alpha*k1)
		xCur = xCur + h
		yCur = yCur + (1.0-1.0/(2.0*alpha))*k1 + (1.0/(2.0*alpha))*k2

		fmt.Printf("y(%f) = %f\n", xCur, yCur)
	}
	return xCur, yCur
}

func RungeKutta2(f func(a float64, b float64) float64, x0 float64, y0 float64, h float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var k1 float64
	var k2 float64
	for range step {
		k1 = f(xCur, yCur)
		k2 = f(xCur+h, yCur+k1)
		xCur = xCur + h
		yCur = yCur + h/2.0*k1 + h/2.0*k2

		fmt.Printf("y(%f) = %f\n", xCur, yCur)
	}
	return xCur, yCur
}

func RungeKutta3(f func(a float64, b float64) float64, x0 float64, y0 float64, h float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var k1 float64
	var k2 float64
	var k3 float64
	for range step {
		k1 = f(xCur, yCur)
		k2 = f(xCur+h/2, yCur+(k1*h)/2)
		k3 = f(xCur+h, yCur-(k1*h)+2*k2*h)
		xCur = xCur + h
		yCur = yCur + h/6.0*(k1+4.0*k2+k3)

		fmt.Printf("y(%f) = %f\n", xCur, yCur)
	}
	return xCur, yCur
}

func RungeKutta4(f func(a float64, b float64) float64, x0 float64, y0 float64, h float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var k1 float64
	var k2 float64
	var k3 float64
	var k4 float64
	for range step {
		k1 = f(xCur, yCur)
		k2 = f(xCur+h/2, yCur+(k1*h)/2)
		k3 = f(xCur+h/2, yCur+(k2*h)/2)
		k4 = f(xCur+h, yCur+k3*h)
		xCur = xCur + h
		yCur = yCur + h/6.0*(k1+2.0*k2+2.0*k3+k4)

		fmt.Printf("y(%f) = %f\n", xCur, yCur)
	}
	return xCur, yCur
}
