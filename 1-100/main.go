package main

import (
	"fmt"
)

type FuncInt func() int

func main() {
    Run("Problem 1: Multiples of 3 or 5", "Sum for 1000: %d", func() int { return MultiplesOf3Or5(1000) })

    Run("Problem 2: Even Fibonacci Numbers", "Sum up to 4 million: %d", func() int { return EvenFibonacciNumbers(4000000) })
}

func Run(title, description string, f FuncInt) {
    fmt.Println(title)
    fmt.Printf(description, f())
    fmt.Print("\n\n")
}
