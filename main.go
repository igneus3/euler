package main

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"

	"euler/globals"
    _"euler/problems_1_100"
)

const usage = `Usage of euler:
    euler --problem [PROBLEM NUMBER] [ARGUMENTS]
Options:
    -p, --problem       Required: Problem number to run
`

func main() {
    var problemNumber int 
    flag.IntVar(&problemNumber, "p", 0, "Problem number to run.")
    flag.IntVar(&problemNumber, "problem", 0, "Problem number to run.")
    flag.Usage = func() { fmt.Print(usage) }

    flag.Parse()

    if problemNumber == 0 {
        panic("Missing problem number!")
    }

    args := flag.Args()
    params := []int{}
    for _, arg := range args {
        param, err := strconv.Atoi(arg)
        if err != nil {
            panic(fmt.Sprintf("Invalid parameter %s", args))
        }
        params = append(params, param)
    }
   
    reflectionItem, ok := globals.ReflectionLookup[problemNumber]
    if !ok {
        panic(fmt.Sprintf("Solution for problem %d not found!", problemNumber))
    }

    result := reflect.ValueOf(reflectionItem.Function).Call([]reflect.Value{reflect.ValueOf(params)})

    fmt.Printf("%s: %v\n", reflectionItem.Title, result[0])
}
