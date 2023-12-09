package main

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
