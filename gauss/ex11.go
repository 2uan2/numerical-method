package gauss

import (
	"fmt"
	"math"
	"os"
)

// Hàm thực hiện giải và ghi kết quả vào file
func solve(f *os.File, caseName string, coeffX3 float64) {
	fmt.Fprintf(f, "\n--- Kết quả câu %s (Hệ số x3: %.1f) ---\n", caseName, coeffX3)

	x1, x2, x3 := 0.0, 0.0, 0.0
	tolerance := 1e-2
	maxIterations := 300

	fmt.Fprintf(f, "%-5s | %-12s | %-12s | %-12s | %-12s\n", "Step", "x1", "x2", "x3", "Error")
	fmt.Fprintln(f, "---------------------------------------------------------------------------")

	for k := 1; k <= maxIterations; k++ {
		oldX1, oldX2, oldX3 := x1, x2, x3

		x1 = 0.2 - (coeffX3 * oldX3)
		x2 = -1.425 + 0.5*x1 + 0.25*oldX3
		x3 = 2.0 - x1 + 0.5*x2

		// Tính sai số
		maxErr := math.Max(math.Abs(x1-oldX1), math.Max(math.Abs(x2-oldX2), math.Abs(x3-oldX3)))

		// Ghi vào file thay vì in ra màn hình
		fmt.Fprintf(f, "%-5d | %-12.4e | %-12.4e | %-12.4e | %-12.4e\n", k, x1, x2, x3, maxErr)

		if math.IsInf(x1, 0) || math.IsNaN(x1) {
			fmt.Fprintf(f, ">> Cảnh báo: Tràn số tại bước %d. Dừng lặp.\n", k)
			break
		}

		if maxErr < tolerance {
			fmt.Fprintf(f, ">> Thành công: Hội tụ đạt sai số < %.1e tại bước %d.\n", tolerance, k)
			break
		}

		if k == maxIterations {
			fmt.Fprintln(f, ">> Thông báo: Đã chạy hết 300 bước mà chưa đạt sai số yêu cầu.")
		}
	}
	fmt.Fprintf(f, "--- Kết thúc câu %s ---\n", caseName)
}

func main() {
	fileName := "ket_qua_gauss_seidel.txt"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Lỗi khi tạo file:", err)
		return
	}
	defer f.Close()

	solve(f, "c", -1.0)

	fmt.Fprintln(f, "\n"+"===========================================================================")

	solve(f, "d", -2.0)
}

