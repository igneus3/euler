package main

import (
	"math"
)

func LargestPalindromeProduct(digits int) int {
    maxDigitNumber := int(math.Pow(10, float64(digits)) - 1)

    result := 1
    for i := maxDigitNumber; i > 0; i-- {
        for j := i; j > 0; j -- {
            product := i * j
            if IsPalindromic(product) {
                result = max(result, product)
            }
        }
    }

    return result
}

func IsPalindromic(number int) bool {
    digitCount := DigitCount(number)
    stack := Stack{}
    prestack := number
    for i := 1; i <= digitCount; i++ {
        stack.Push(prestack % 10)
        prestack /= 10
    }

    poststack := number
    for stack.Count() > 0 {
        if poststack % 10 != stack.Pop() {
            return false
        }
        poststack /= 10
    }

    return true
}

func DigitCount(number int) int {
    if number == 0 {
        return 1
    }

    result := 0
    for number != 0 {
        number /= 10
        result++
    }

    return result
}

type Node struct {
    Val int
    Next *Node
}

type Stack struct {
    head *Node
    length int
}

func (s *Stack) Push(item int) {
    s.length++

    newNode := &Node{Val: item, Next: s.head}

    s.head = newNode
}

func (s *Stack) Pop() int {
    if s.length == 0 {
        return 0
    }

    s.length--

    node := s.head
    s.head = node.Next
    node.Next = nil

    return node.Val
}

func (s *Stack) Peek() int {
    if s.length == 0 {
        return 0
    }

    return s.head.Val
}

func (s *Stack) Count() int {
    return s.length
}
