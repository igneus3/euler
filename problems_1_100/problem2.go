package problems_1_100

import "euler/globals"

func init() {
    globals.ReflectionLookup[2] = globals.ReflectionItem{Title: "Even Fibonacci Numbers", Function: Problem2}
}

func Problem2(params []int) int {
    return EvenFibonacciNumbers(params[0])
}

func EvenFibonacciNumbers(number int) int {
    prevOne, prevTwo := 1, 1
    fibonacci := 0
    result := 0

    for fibonacci <= number {
       fibonacci = prevOne + prevTwo 
       prevTwo = prevOne
       prevOne = fibonacci

       if fibonacci % 2 == 0 {
           result += fibonacci
       }
    }

    return result
}
