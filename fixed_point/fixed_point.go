package fixed_point

import (
	"errors"
	"fmt"
	"math"
)

func Fixed_point(p_old float64, n0 int, TOL float64, f func(float64) float64) (float64, error) {
	var p_new float64
	for i := 0; i < n0; i++ {
		p_new = f(p_old)
		fmt.Printf("(%f, %f)\n", p_old, p_new)
		if math.Abs(p_new-p_old) < TOL {
			return p_new, nil
		}
		p_old = p_new
	}
	return 0, errors.New("reached max iterations")
}

func round(val float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(val*shift) / shift
}
