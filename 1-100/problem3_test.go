package main

import (
	"fmt"
	"testing"
)

func TestLargestPrimeFactor(t *testing.T) {
    tests := []struct{
        intput int
        output int
    }{
        {2, 2},
        {3, 3},
        {5, 5},
        {7, 7},
        {8, 2},
        {6857, 6857},
        {13195, 29},
        {600851475143, 6857},
    }

    for _, test := range tests {
        test := test
        t.Run(fmt.Sprint(test.intput), func(t *testing.T) {
            if got, want := LargestPrimeFactor(test.intput), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}

func TestIsPrime(t *testing.T) {
    tests := []struct{
        intput int
        output bool
    }{
        {1, false},
        {2, true},
        {3, true},
        {4, false},
        {5, true},
        {7, true},
        {8, false},
        {199, true},
        {200, false},
        {313, true},
        {314, false},
    }

    for _, test := range tests {
        test := test
        t.Run(fmt.Sprint(test.intput), func(t *testing.T) {
            if got, want := IsPrime(test.intput), test.output; got != want {
                t.Fatalf("got = %t, want = %t", got, want)
            }
        })
    }
}
