package root

import (
	"errors"
	"fmt"
	"math"
)

func secant(p0 float64, n0 int, TOL float64, p1 float64, f func(float64) float64) (float64, error) {
	var p float64
	var q0 = f(p0)
	var q1 = f(p1)
	for i := 0; i < n0; i++ {
		p = p1 - q1*(p1-p0)/(q1-q0)
		fmt.Printf("(%f, %f)\n", p0, p)
		if math.Abs(p-p0) < TOL {
			return p, nil
		}
		p0 = p1
		q0 = q1
		p1 = p
		q1 = f(p)
	}
	return 0, errors.New("reached max iterations")
}
