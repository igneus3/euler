package problems_1_100

import (
	"euler/globals"
	"math"
)

func init() {
    globals.ReflectionLookup[9] = globals.ReflectionItem{Title: "Special Pythagorean Triplet", Function: Problem9}
}

func Problem9(params []int) int {
    return SpecialPythagoreanTriplet(params[0])
}

func SpecialPythagoreanTriplet(x int) int {
    for a := 1; a <= x; a++ {
        for b := a + 1; b <= x; b++ {
            c := x - a - b
            if c > b {
                if math.Pow(float64(a), 2) + math.Pow(float64(b), 2) == math.Pow(float64(c), 2) {
                    return a * b * c
                }
            }
        }
    }

    return 0
}

