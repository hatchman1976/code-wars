package kata

import "math"
//import "fmt"


func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func Step(g, m, n int) []int {

	for ; m < n; m++ {

		if IsPrime(m) {
			for i := m + 1; i - m <= g; i++ {
				if IsPrime(i) && i - m == g {
					//Pairs have been found
					return []int {m, i}
				}
			}
		}
	}
	return nil
}
