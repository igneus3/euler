package problems_1_100

import "euler/globals"

func init() {
    globals.ReflectionLookup[5] = globals.ReflectionItem{Title: "Smallest Multiple", Function: Problem5}
}

func Problem5(params []int) int {
    return SmallestMultiple(params[0])
}

func SmallestMultiple(number int) int {
    if number == 1 {
        return 1
    }

    current := SmallestMultiple(number - 1)
    return number * current / GreatestCommonDenominator(number, current)
}

func GreatestCommonDenominator(a, b int) int {
    x, y := a, b
    if x < y {
        x, y = b, a
    }

    r := x % y
    if r == 0 {
        return y
    }

    return GreatestCommonDenominator(y, r)
}

