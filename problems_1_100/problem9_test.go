package problems_1_100

import (
	"fmt"
	"testing"
)

func TestSpecialPythagoreanTriplet(tt *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {12, 60},
        {1000, 31875000},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := SpecialPythagoreanTriplet(test.input), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}

