package main

import (
  "fmt"
  "math/rand"
  "time"
)


func quicksort(numbers []int) {
    if len(numbers) <= 1 {
        return
    }

    lastElement := uint(len(numbers) - 1)

    pivotVal := numbers[lastElement];

    lastSwapped := uint(0);
    for k, _ := range numbers[:lastElement] {
        if numbers[k] <= pivotVal {
            tempVal := numbers[k]
            numbers[k] = numbers[lastSwapped]
            numbers[lastSwapped] = tempVal

            lastSwapped++
        }
    }

    tempVal := numbers[lastSwapped]
    numbers[lastSwapped] = numbers[lastElement]
    numbers[lastElement] = tempVal

    quicksort(numbers[:lastSwapped])
    quicksort(numbers[lastSwapped:])
}


func bubblesort(numbers []int) {
    lastElement := len(numbers) - 1

    for i := range numbers[:lastElement] {
        for j := range numbers[i:lastElement] {
            if numbers[j] > numbers[j+1] {
                tempVal := numbers[j]
                numbers[j] = numbers[j+1]
                numbers[j+1] = tempVal
            }
        }
    }
}


func verifysorted(numbers []int) bool {
    lastElement := len(numbers) - 1

    for i := range numbers[:lastElement] {
        if numbers[i] > numbers[i+1] {
          return false
        }
    }

    return true
}



func trace(f func()) time.Duration {
    startTime := time.Now()

    f()

    return time.Now().Sub(startTime)
}

func main() {
    N := 100000

    numbers1 := make([]int, N)
    numbers2 := make([]int, N)

    for i, _ := range numbers1 {
      random := rand.Int()
      numbers1[i] = random
      numbers2[i] = random
    }

    b1 := trace(func() {
        quicksort(numbers1)
    })
    fmt.Printf("quicksort:  %s\n", b1)

    b2 := trace(func() {
        bubblesort(numbers2)
    })
    fmt.Printf("bubblesort: %s\n", b2)
}
