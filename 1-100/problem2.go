package main

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
