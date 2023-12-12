package main

import (
	"fmt"
	"testing"
)

func TestLargestPalindromeProduct(tt *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {1, 9},
        {2, 9009},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func (t *testing.T) {
            if got, want := LargestPalindromeProduct(test.input), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}

func TestIsPalindromic(tt *testing.T) {
    tests := []struct{
        input int
        output bool
    }{
        {9, true},
        {9009, true},
        {123454321, true},
        {1234, false},
        {988789, false},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func (t *testing.T) {
            if got, want := IsPalindromic(test.input), test.output; got != want {
                t.Fatalf("got = %t, want = %t", got, want)
            }
        })
    }
}
