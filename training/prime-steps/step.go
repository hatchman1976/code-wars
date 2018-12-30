package kata

import "math"

func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func Step(g, m, n int) []int {

	for ; m < n; m++ {

		if m%2 != 0 {
			if (m + g) <= n {
				if IsPrimeSqrt(m) && IsPrimeSqrt(m+g) {
					return []int{m, m + g}
				}
			}
		}
	}
	return nil
}
