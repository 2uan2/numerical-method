package interpolation

import "fmt"

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

func GeneralLagrange(x float64, xs []float64, ys []float64) float64 {
	p := 0.0
	for idx, cur_x := range xs {
		p += (l(idx, x, xs) / l(idx, cur_x, xs)) * ys[idx]
		fmt.Printf("L%d = %f\n", idx, l(idx, x, xs)/l(idx, cur_x, xs))
	}
	return p
}

func l(k int, x float64, others []float64) float64 {
	l := 1.0
	for idx, cur_x := range others {
		if idx == k {
			continue
		}

		l *= (x - cur_x)
		// fmt.Printf("l *= (%f - %f)\n", x, cur_x)
	}
	fmt.Printf("l%d(%f) = %f\n", k, x, l)
	return l
}
