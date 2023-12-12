package main

import (
	"fmt"
)

type FuncInt func() int

func main() {
    Run("Problem 1: Multiples of 3 or 5", "Sum for 1000: %d", func() int { return MultiplesOf3Or5(1000) })

    Run("Problem 2: Even Fibonacci Numbers", "Sum up to 4 million: %d", func() int { return EvenFibonacciNumbers(4000000) })

    Run("Problem 3: Largest Prime Factor", "Largest Prime Factor for 600851475143: %d", func() int { return LargestPrimeFactor(600851475143) })

    Run("Problem 4: Largest Palindrome Product", "Largest Palindrome for 3 digit product: %d", func() int { return LargestPalindromeProduct(3) })
}

func Run(title, description string, f FuncInt) {
    fmt.Println(title)
    fmt.Printf(description, f())
    fmt.Print("\n\n")
}
