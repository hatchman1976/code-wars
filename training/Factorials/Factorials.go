package kata

import (
	"fmt"
	//"strconv"

	"math/big"
	"sync"

	"github.com/shopspring/decimal"
)

func CalculateFactorial(i int64) decimal.Decimal {

	if i <= 0 {
		returnVal, _ := decimal.NewFromString("1")
		return returnVal
	}
	returnVal := CalculateFactorial(i - 1)
	
	return  decimal.New(i, 0).Mul(returnVal)
}
func Going(n int) float64 {
	// your code
	results := make([]decimal.Decimal, n)

	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 1; i <= n; i++ {
		go func(i int64) {
			defer wg.Done()
			returnVal := CalculateFactorial(i)
			results[i-1] = returnVal
		}(int64(i))
	}
	wg.Wait()

	fmt.Println(results)

	calculatedFact := new(big.Int)
	for i, j := range results {

	}

}
