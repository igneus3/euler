package problems_1_100

import "euler/globals"

func init() {
    globals.ReflectionLookup[1] = globals.ReflectionItem{Title: "Multiples of 3 and 5", Function: Problem1}
}

func Problem1(params []int) int {
    return MultiplesOf3Or5(params[0])
}

func MultiplesOf3Or5(number int) int {
    result := 0
    number--
    for number > 0 {
        if number % 5 == 0 {
            result += number
        } else if number % 3 == 0 {
            result += number
        }
        number--
    }

    return result
}
