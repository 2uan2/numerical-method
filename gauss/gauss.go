package gauss

import (
	"fmt"
	"math"
)

func Gauss(A0 [][]int, b []int) []int {
	n := len(A0)
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			m := A0[j][i] / A0[i][i]
			for k := range n {
				A0[j][k] = A0[j][k] - m*A0[i][k]
			}
			b[j] = b[j] - m*b[i]
		}
	}

	x := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		right := b[i]
		for j := i + 1; j < n; j++ {
			right -= A0[i][j] * x[j]
		}
		x[i] = right / A0[i][i]
	}
	return x
}

func GaussSeidel(A [][]float64, b []float64, x []float64, iterations int, threshold float64) []float64 {
	n := len(A)
	newA := make([][]float64, n)
	for i := range n {
		newA[i] = make([]float64, n+1)
	}

	for i := range n {
		for j := range n {
			if i == j {
				continue
			}
			newA[i][j] = -A[i][j]
		}
		newA[i][n] = b[i]
	}
	fmt.Println(newA)

	fmt.Println(x)
	for i := range iterations {
		var max_diff float64 = 0
		// loop for lines
		for j := range newA {
			var sum float64 = 0
			// loop for columns
			for k := range n {
				sum += newA[j][k] * x[k]
			}
			sum += newA[j][n]
			if math.Abs(sum-x[j]) >= max_diff {
				max_diff = math.Abs(sum - x[j])
			}
			x[j] = sum
		}

		fmt.Printf("[%d]: %v\n", i, x)
		if max_diff <= threshold {
			break
		}
	}

	return x
}

func IsDiagonallyDominant(A [][]float64) bool {
	n := len(A)
	diagonallyDominant := true
	for i := range n {
		var sum float64 = 0
		for j := range n {
			if i == j {
				continue
			}
			sum += A[i][j]
		}
		if A[i][i] <= math.Abs(sum) {
			diagonallyDominant = false
		}
	}
	return diagonallyDominant
}

// [[2 1 1 0]  [6
//  [0 1 1 1]   3
//  [0 0 2 2]   8
//  [0 0 0 2]]  2]

// [[2 1 1 0]  [6
//  [4 3 3 1]   15
//  [8 7 9 5]   41
//  [6 7 9 8]]  40]
