package kata

import (
	"math"
)

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {

	var month int
	var savings float64

	startPriceNewDecrease := float64(startPriceNew)
	startPriceOldDecrease := float64(startPriceOld)

	if startPriceOld < startPriceNew {
		for savings <= startPriceNewDecrease-startPriceOldDecrease {

			month++
			if month%2 == 0 {
				percentLossByMonth += 0.5
			}

			savings += float64(savingperMonth)
			startPriceOldDecrease -= startPriceOldDecrease * percentLossByMonth / 100
			startPriceNewDecrease -= startPriceNewDecrease * percentLossByMonth / 100
		}
	}
	return [2]int{month, int(math.Round(savings - ((startPriceNewDecrease) - startPriceOldDecrease)))}
}
