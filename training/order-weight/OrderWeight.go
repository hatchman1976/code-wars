package kata

import (
	"sort"
	"strconv"
	"strings"
)

func OrderWeight(strng string) string {

	type weight struct {
		actual     string
		calculated int
	}
	var weights []weight

	newweights := strings.Fields(strng)

	for t := range newweights {
		var scaled int
		for _, z := range strings.Split(newweights[t], "") {
			var val2, _ = strconv.Atoi(z)
			scaled += val2
		}
		weights = append(weights, weight{newweights[t], scaled})
	}

	sort.Slice(weights, func(i, j int) bool {
		if weights[i].calculated < weights[j].calculated {
			return true
		}
		if weights[i].calculated > weights[j].calculated {
			return false
		}
		return weights[i].actual < weights[j].actual
	})

	returnValue := ""
	for _, y := range weights {
		returnValue += y.actual + " "
	}
	return strings.Trim(returnValue, " ")
}
