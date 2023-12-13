package problems_1_100

import (
	"euler/globals"
	"math"
)

func init() {
    globals.ReflectionLookup[6] = globals.ReflectionItem{Title: "Sum Square Difference", Function: Problem6}
}

func Problem6(params []int) int {
    return SumSquareDifference(params[0])
}

func SumSquareDifference(x int) int {
    return squareOfSum(x) - sumOfSquares(x)
}

func sumOfSquares(x int) int {
    sum := 0.0
    i := 1.0
    for ; i <= float64(x); i++ {
        sum += math.Pow(i, 2)
    }
    return int(sum)
}

func squareOfSum(x int) int {
    sum := 0
    for i := 1; i <= x; i++ {
        sum += i
    }
    return int(math.Pow(float64(sum), 2))
}
