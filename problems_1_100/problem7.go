package problems_1_100

import (
	"euler/globals"
)

func init() {
    globals.ReflectionLookup[7] = globals.ReflectionItem{Title: "Nth Prime Number", Function: Problem7}
}

func Problem7(params []int) int {
    return NthPrimeNumber(params[0])
}

func NthPrimeNumber(x int) int {
    primes := primeLinkedList{}

    for i := 1; i <= x; i++ {
        primes.GenerateNextPrime()
    }

    return primes.LastPrime()
}

type primeNode struct {
    Value int
    Next *primeNode
}

type primeLinkedList struct {
    head *primeNode
    tail *primeNode
    length int
}

func (p *primeLinkedList) GenerateNextPrime() int {
    if p.length < 1 || p.tail == nil {
        p.length++
        newNode := &primeNode{Value: 2}
        p.head = newNode
        p.tail = newNode
    } else {
        newNode := &primeNode{Value: p.tail.Value + 1}
        for !p.testPossiblePrime(newNode) {
            newNode.Value++
        }

        p.length++
        p.tail.Next = newNode
        p.tail = p.tail.Next
    }

    return p.tail.Value
}

func (p *primeLinkedList) testPossiblePrime(node *primeNode) bool {
    cur := p.head 
    for cur != nil {
        if node.Value % cur.Value == 0 {
            return false
        }
        cur = cur.Next
    }

    return true
}

func (p *primeLinkedList) LastPrime() int {
    return p.tail.Value
}

func (p *primeLinkedList) Count() int {
    return p.length
}
