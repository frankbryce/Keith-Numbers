package main

import (
    "fmt"
    "keith"
)

func main() {
    lt := 100000
    keiths := []int{}
    for i := 10; i <= lt; i++ {
        if keith.IsKeith(i) {
            keiths = append(keiths, i)
        }
    }
    fmt.Println(keiths)
}