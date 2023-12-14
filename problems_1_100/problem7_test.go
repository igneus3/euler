package problems_1_100

import (
	"fmt"
	"testing"
)

func TestNthPrimeNumber(tt *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {1, 2},
        {6, 13},
        {10, 29},
        {100, 541},
        {1000, 7919},
        {10001, 104743},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := NthPrimeNumber(test.input), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}

