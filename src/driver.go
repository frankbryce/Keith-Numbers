package main

import (
    "fmt"
    "keith"
    "math/big"
)

func main() {
    keiths := []*big.Int{}
    z := big.NewInt(0)
    inc := big.NewInt(int64(1))
    upto := big.NewInt(int64(1 << 25))
    for i := big.NewInt(10); i.Cmp(upto) <= 0; i.Add(i, inc) {
        if keith.IsKeith(i) {
            keiths = append(keiths, big.NewInt(0).Add(z, i))
        }
    }
    fmt.Println(keiths)
    fmt.Println(keith.IsKeith(big.NewInt(int64(1 << 62))))
}