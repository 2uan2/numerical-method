package ode

import "fmt"

func Midpoint(f func(a float64, b float64) float64, x0 float64, y0 float64, h float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var xMid float64
	var yMid float64
	// fmt.Println("xCur: ", xCur)
	// fmt.Println("yCur: ", yCur)
	// fmt.Println("====")

	for range step {
		yMid = yCur + (h/2)*f(xCur, yCur)
		xMid = xCur + h/2
		// fmt.Println("xMid: ", xMid)
		// fmt.Println("yMid: ", yMid)
		yCur = yCur + h*f(xMid, yMid)
		xCur = xCur + h
		fmt.Printf("y(%f) = %f\n", xCur, yCur)
		// fmt.Println("xCur: ", xCur)
		// fmt.Println("yCur: ", yCur)
		// fmt.Println("====")
	}

	return xCur, yCur
}

func Heun(f func(a float64, b float64) float64, x0 float64, y0 float64, h float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var yPred float64
	var yCor float64
	for range step {
		yPred = yCur + h*f(xCur, yCur)
		yCor = yCur + h/2*(f(xCur, yCur)+f(xCur+h, yPred))
		xCur = xCur + h
		yCur = yCor

		// fmt.Println("xCur: ", xCur)
		// fmt.Println("yCur: ", yCur)
		fmt.Printf("y(%f) = %f\n", xCur, yCur)
		fmt.Println("====")
	}
	return xCur, yCur
}

func RungeKutta2(f func(a float64, b float64) float64, x0 float64, y0 float64, h float64, alpha float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var k1 float64
	var k2 float64
	for range step {
		k1 = h * f(xCur, yCur)
		k2 = h * f(xCur+alpha*h, yCur+alpha*k1)
		// k2 = yCur + h/2*(f(xCur, yCur)+f(xCur+h, k1))
		xCur = xCur + h
		yCur = yCur + (1.0-1.0/(2.0*alpha))*k1 + (1.0/(2.0*alpha))*k2

		// fmt.Println("xCur: ", xCur)
		// fmt.Println("yCur: ", yCur)
		fmt.Printf("y(%f) = %f\n", xCur, yCur)
		fmt.Println("====")
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
		// k2 = yCur + h/2*(f(xCur, yCur)+f(xCur+h, k1))
		xCur = xCur + h
		yCur = yCur + h/6.0*(k1+4.0*k2+k3)

		fmt.Printf("y(%f) = %f\n", xCur, yCur)
		fmt.Println("====")
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
		fmt.Println("====")
	}
	return xCur, yCur
}

// func RungeKuttaSystem(fs []func(args ...float64) float64, x float64, initials []float64, h float64, step int) (float64, float64) {
// 	xCur := x
// 	zCur := initials
// 	zNew := make([]float64, len(initials))
// 	for range step {
// 		f := fs[0]
// 		zNew[0] = zCur[0] + h*f(initials...)
//
// 	}
//
// 	return 1.0, 2.0
// }
