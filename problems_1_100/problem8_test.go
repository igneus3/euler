package problems_1_100

import (
	"fmt"
	"testing"
)

func TestLargestProductInASeries(tt *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {4, 5832},
        {13, 23514624000},
    }

    for _, test := range tests {
        test := test
        tt.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := LargestProductInASeries(test.input), test.output; got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}

