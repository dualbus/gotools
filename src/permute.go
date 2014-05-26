package main

// Heap's algorithm
// <http://en.wikipedia.org/wiki/Heap's_algorithm>

import (
    "fmt"
    "flag"
)

func permute(length int, data []rune) {
    if length == 1 {
        fmt.Println(string(data))
        return
    } else {
        for c := 0; c < length; c++ {
            var k int

            permute(length - 1, data)

            if length % 2 == 0 {
                k = c
            } else {
                k = 0
            }

            tempVal := data[k]
            data[k] = data[length - 1]
            data[length - 1] = tempVal
        }
    }
}

func main() {
    flag.Parse()

    if flag.NArg() < 1 {
        return
    }

    runes := []rune(flag.Arg(0))

    permute(len(runes), runes)
}
