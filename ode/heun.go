package ode

import "fmt"

func Heun(f func(a, b float64) float64, x0, y0, h float64, step int) (float64, float64) {
	xCur := x0
	yCur := y0
	var yPred float64
	var yCor float64
	for range step {
		yPred = yCur + h*f(xCur, yCur)
		yCor = yCur + h/2*(f(xCur, yCur)+f(xCur+h, yPred))
		xCur = xCur + h
		yCur = yCor

		fmt.Printf("y(%f) = %f\n", xCur, yCur)
	}
	return xCur, yCur
}

// k: amount of correction step
func HeunMultistep(f func(a, b float64) float64, x0, y0, x1, y1, h float64, step int, k int) (float64, float64) {
	// xPrev would be x_(i-1) and xCur would be x_i so that the corrector and predictor can be x_(i+1)
	xPrev := x0
	yPrev := y0
	xCur := x1
	yCur := y1
	var yPred float64
	var yCor float64
	for range step {
		yPred = yCur + h*(3.0/2.0*f(xCur, yCur)-1.0/2.0*f(xPrev, yPrev))

		// set yCor to yPred initially to start loop, yPred is y_(x+1)[0th iteration]
		yCor = yPred
		// do k correction step
		for range k {
			yCor = yCur + h*(1.0/2.0*f(xCur+h, yCor)+1.0/2.0*f(xCur, yCur))
		}
		xPrev = xCur
		yPrev = yCur
		xCur = xCur + h
		yCur = yCor
		// fmt.Printf("y(%f) = %f\n", xCur, yCur)
	}
	return xCur, yCur
}

// HeunMultistep3Step solves ODE using 3 historical points.
// You must provide 3 initialization points: (x0,y0), (x1,y1), (x2,y2)
func HeunMultistep3Step(f func(a, b float64) float64, x0, y0, x1, y1, x2, y2, h float64, step int, k int) (float64, float64) {
	// Initialize history chains
	// index 2 is current (i), index 1 is (i-1), index 0 is (i-2)
	xHist := []float64{x0, x1, x2}
	yHist := []float64{y0, y1, y2}

	var yPred float64
	var yCor float64

	for i := 0; i < step; i++ {
		// Extract historical terms to keep code readable and matching your slides
		fI := f(xHist[2], yHist[2])     // f(x_i, y_i)
		fPrev1 := f(xHist[1], yHist[1]) // f(x_{i-1}, y_{i-1})
		fPrev2 := f(xHist[0], yHist[0]) // f(x_{i-2}, y_{i-2})

		// 3-Step Predictor formula from slide
		yPred = yHist[2] + h*((23.0/12.0)*fI-(16.0/12.0)*fPrev1+(5.0/12.0)*fPrev2)

		xNext := xHist[2] + h
		yCor = yPred

		// Inner correction loop (k iterations)
		for j := 0; j < k; j++ {
			// 3-Step Corrector formula from slide
			yCor = yHist[2] + h*((5.0/12.0)*f(xNext, yCor)+(8.0/12.0)*fI-(1.0/12.0)*fPrev1)
		}

		// Rotate history window forward across time: drop oldest, push newest
		xHist[0], xHist[1], xHist[2] = xHist[1], xHist[2], xNext
		yHist[0], yHist[1], yHist[2] = yHist[1], yHist[2], yCor

		fmt.Printf("y(%f) = %f\n", xHist[2], yHist[2])
	}

	return xHist[2], yHist[2]
}

// HeunMultistep4Step solves ODE using 4 historical points.
// You must provide 4 initialization points: (x0,y0) through (x3,y3)
func HeunMultistep4Step(f func(a, b float64) float64, x0, y0, x1, y1, x2, y2, x3, y3, h float64, step int, k int) (float64, float64) {
	// Initialize history chains
	// index 3 is current (i), 2 is (i-1), 1 is (i-2), 0 is (i-3)
	xHist := []float64{x0, x1, x2, x3}
	yHist := []float64{y0, y1, y2, y3}

	var yPred float64
	var yCor float64

	for i := 0; i < step; i++ {
		fI := f(xHist[3], yHist[3])     // f(x_i, y_i)
		fPrev1 := f(xHist[2], yHist[2]) // f(x_{i-1}, y_{i-1})
		fPrev2 := f(xHist[1], yHist[1]) // f(x_{i-2}, y_{i-2})
		fPrev3 := f(xHist[0], yHist[0]) // f(x_{i-3}, y_{i-3})

		// 4-Step Predictor formula from your new slide
		yPred = yHist[3] + (h/24.0)*(55.0*fI-59.0*fPrev1+37.0*fPrev2-9.0*fPrev3)

		xNext := xHist[3] + h
		yCor = yPred

		// Inner correction loop (k iterations)
		for j := 0; j < k; j++ {
			// 4-Step Corrector formula from your new slide
			yCor = yHist[3] + (h/24.0)*(9.0*f(xNext, yCor)+19.0*fI-5.0*fPrev1+fPrev2)
		}

		// Rotate history window forward
		xHist[0], xHist[1], xHist[2], xHist[3] = xHist[1], xHist[2], xHist[3], xNext
		yHist[0], yHist[1], yHist[2], yHist[3] = yHist[1], yHist[2], yHist[3], yCor

		fmt.Printf("y(%f) = %f\n", xHist[3], yHist[3])
	}

	return xHist[3], yHist[3]
}
