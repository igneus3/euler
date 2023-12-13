package problems_1_100

import (
	"fmt"
	"testing"
)

func TestSumSquareDifference(tt *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {10, 2640},
        {20, 41230},
        {100, 25164150},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := SumSquareDifference(test.input), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}

func TestSumOfSquares(tt *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {10, 385},
        {20, 2870},
        {100, 338350},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := sumOfSquares(test.input), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }

}

func TestSquareOfSum(tt *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {10, 3025},
        {20, 44100},
        {100, 25502500},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := squareOfSum(test.input), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }

}

