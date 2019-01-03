package kata

import "sort"
import "strings"
import "strconv"
import "fmt"

func OrderWeight(strng string) string {

	weights := strings.Fields(strng)

	weightOrder := [2][]int{}

	weightOrder[0] = make([]int, len(weights))
	weightOrder[1] = make([]int, len(weights))

	fmt.Println(weightOrder)

	for t := range weights {
		val, _ := strconv.Atoi(weights[t])
		weightOrder[0][t] = val
		for _, z := range strings.Split(weights[t], "") {
			val, _ = strconv.Atoi(z)
			weightOrder[1][t] += val
		}
	}
	fmt.Println(weightOrder)

	newweights := strings.Fields(strng)

	maps := make(map[string]int)

	keys := make([]string, 0, len(newweights))

	for t := range newweights {
		val, _ := strconv.Atoi(newweights[t])
		var scaled int
		for _, z := range strings.Split(newweights[t], "") {
			var val2, _ = strconv.Atoi(z)
			scaled += val2
		}

		keys = append(keys)
	}

	matrix := [2][5]int{
		{9, 1, 3, 5, 6},
		{4, 5, 6, 7, 8},
	}
	sort.Slice(matrix[:], func(i, j int) bool {
		for x := range matrix[i] {
			fmt.Println(matrix[i][x])
			fmt.Println(matrix[j][x])
			if matrix[i][x] == matrix[j][x] {

				continue
			}
			return matrix[i][x] < matrix[j][x]
		}
		return false
	})

	fmt.Println(matrix)

	return "something"
}
