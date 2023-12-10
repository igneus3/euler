package main

func LargestPrimeFactor(number int) int {
    if IsPrime(number) {
        return number
    }

    i := 2
    largest := 2
    for ; i <= number; i++ {
        if !IsPrime(i) {
            continue
        }

        for number % i == 0 {
            if i == 2 {
                number >>= 1
            } else {
                number /= i
            }
        }

        largest = i
    }

    return largest
}

func IsPrime(number int) bool {
    if number == 2 || number == 3 {
        return true
    }

    if number <= 1 || number % 2 == 0 || number % 3 == 0 {
        return false
    }

    for i := 5; i * i <= number; i += 6 {
        if number % i == 0 || number % (i + 2) == 0 {
            return false
        }
    }

    return true
}
