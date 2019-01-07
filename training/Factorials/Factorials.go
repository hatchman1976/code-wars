package kata

import (
	"fmt"
	//"strconv"
	"math"
	"math/big"
	"sync"
)

func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

func CalculateFactorial(i float64) *big.Float {

	if i <= 0 {
		return big.NewFloat(1)
	}
	returnVal := CalculateFactorial(i - 1)

	return new(big.Float).Mul(big.NewFloat(i), returnVal)
}
func Going(n int) float64 {
	// your code
	results := make([]big.Float, n)

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i float64) {
			defer wg.Done()
			returnVal := CalculateFactorial(i)
			results[int(i)-1] = *returnVal
		}(float64(i))
	}
	wg.Wait()

	fmt.Println(results)

	floatReturn := new(big.Float)
	for _, j := range results {
		floatReturn.Add(floatReturn, &j)
	}

	floatCalc := new(big.Float)
	floatCalc.Quo(big.NewFloat(1), &results[n-1])
	floatReturn.Mul(floatReturn, floatCalc)
	floatReturn.SetPrec(10)
	fmt.Println(floatReturn)

valr, _ := floatReturn.Float64()
	return valr

}
