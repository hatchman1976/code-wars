package kata

import "math"



func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func Step(g, m, n int) []int {

	 var primeCount = make([]int, n - m)

	 i := 0
	 for  ; m <= n; m++ {
		 if IsPrime(m) {
			primeCount[i] = m
			i++
		 }
	 }

	 var primeActual = make([]int, i)

	 for i = 0; i < len(primeActual); i++ {
		 primeActual[i] = primeCount[i]
	 }

	 primeCount = nil


	for k := range primeActual {
		for l := range primeActual {
			if primeActual[l] - primeActual[k] == g {
				return []int {primeActual[k], primeActual[l]}
			}
		}
	}

	return nil
}
