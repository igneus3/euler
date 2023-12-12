package main

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

