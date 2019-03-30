package kata

import (
	"math/big"
	"strconv"
)

func Going(n int) float64 {

	factorialSum := new(big.Int).SetInt64(1)
	factorialAdd := new(big.Int)

	for i := 1; i <= n; i++ {

		factorialSum.Mul(factorialSum, new(big.Int).SetInt64(int64(i)))
		factorialAdd.Add(factorialAdd, factorialSum)
	}

	otherRat := big.NewRat(factorialAdd.Int64(), factorialSum.Int64())
	floatanswer, _ := strconv.ParseFloat(otherRat.FloatString(7)[0:8], 64)
	return floatanswer
}
