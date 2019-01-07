package kata

import (
	"fmt"
	//"strconv"
	"math"
	"sync"
	"math/big"
)

func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

func CalculateFactorial(i int64) *big.Int {

	if i <= 0 {
		return big.NewInt(1)
	}
	returnVal := CalculateFactorial(i-1)

	return new(big.Int).Mul(big.NewInt(i), returnVal)
}
func Going(n int) float64 {
	// your code
	results := make([]big.Int, n)

	 wg := sync.WaitGroup{}
	wg.Add(n)
	
	var i int64
	for i := 1; i <= n; i++ {
		go func(i int64) {
			defer wg.Done()
			returnVal := CalculateFactorial(i)
			results[i-1] = *returnVal
		}(int64(i))
	}
	wg.Wait()

	//fmt.Println(results)

	factorialSum := big.NewInt(0)
	for _, j := range results {
		factorialSum.Add(factorialSum, &j)
	}

	fmt.Println(n, factorialSum)

 	factorialSum.Mul(factorialSum.Div(big.NewInt(1), factorialSum), factorialSum)

}
