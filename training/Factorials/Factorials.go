package kata

import (
	"fmt"
	"strconv"

	"math/big"
	"sync"

	//"github.com/shopspring/decimal"
)

func CalculateFactorial(i int64) *big.Int {

	if i <= 0 {
		returnVal := new(big.Int).SetInt64(int64(1))
		return returnVal
	}
	returnVal := CalculateFactorial(i - 1)
	
	return new(big.Int).Mul(big.NewInt(i), returnVal)
}
func Going(n int) float64 {
	// your code
	results := make([]big.Int, n)

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int64) {
			defer wg.Done()
			returnVar := CalculateFactorial(i)
			results[i-1] = *returnVar
		}(int64(i))
	}
	wg.Wait()

	//fmt.Println(results)

	calculatedFact := new(big.Int)
	for _, j := range results {
		calculatedFact.Add(calculatedFact, &j)
	}

	//returnVar, _ :=  new(big.Float).Quo(new(big.Float).SetInt(calculatedFact), new(big.Float).SetInt(&results[n-1])).Float64()
	
	otherRat := big.NewRat(calculatedFact.Int64(), results[n-1].Int64())

	//anotherRat, _ := strconv.ParseFloat(otherRat.FloatString(7), 64)

	secondrat := otherRat.FloatString(7)
	thridRat, _ := strconv.ParseFloat(secondrat[0:len(secondrat)-1], 64)
	
	fmt.Println(thridRat)
	
	return thridRat


}
