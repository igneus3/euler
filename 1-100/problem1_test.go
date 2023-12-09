package main

import (
	"fmt"
	"testing"
)

func TestMultiplesOf3Or5(t *testing.T) {
    tests := []struct{
        input int
        output int
    }{
        {10, 23},
        {49, 543},
        {1000, 233168},
        {8456, 16687353},
        {19564, 89301183},
    }

    for _, test := range tests {
        test := test
        t.Run(fmt.Sprint(test.input), func(t *testing.T) {
            if got, want := MultiplesOf3Or5(test.input), test.output;  got != want {
                t.Fatalf("got = %d, want = %d", got, want)
            }
        })
    }
}
