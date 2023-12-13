package problems_1_100

import (
	"fmt"
	"testing"
)

func TestEvenFibonacciNumbers(t *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {80, 44},
        {89, 188},
        {4000000, 4613732},
    }

    for _, test := range tests {
        test := test
        t.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := EvenFibonacciNumbers(test.input), test.output;  got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}
