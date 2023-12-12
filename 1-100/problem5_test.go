package main

import (
	"fmt"
	"testing"
)

func TestSmallestMultiple(tt *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {3, 6},
        {5, 60},
        {7, 420},
        {10, 2520},
        {13, 360360},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := SmallestMultiple(test.input), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}

func TestGreatestCommonDenominator(tt *testing.T) {
    tests := []struct{
        a int
        b int
        output int
    }{
        {10, 4, 2},
        {100, 40, 20},
        {210, 45, 15},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprintf("a = %d, b = %d", test.a, test.b), func(t *testing.T) {
            if got, want := GreatestCommonDenominator(test.a, test.b), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }

}

