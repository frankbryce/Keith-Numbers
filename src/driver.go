package main

import (
    "fmt"
    "keith"
    "math/big"
    //"time"
)

func main() {
    n := big.NewInt(int64(10000000000000))
    n.Mul(n, n)
    n.Mul(n, big.NewInt(10))
    
    zero := big.NewInt(0)
    one := big.NewInt(1)
    nineteen := big.NewInt(19)
    for mod := big.NewInt(0).Mod(n, nineteen); mod.Cmp(zero) != 0; mod = big.NewInt(0).Mod(n, nineteen) {
        n.Add(n, one)
    }
    
    c := int64(0)
    for ;!keith.IsKeith(n); n.Add(n, nineteen) {
        if c % 100000 == 0 {
            fmt.Println("%v : %v digits", n, len(n.String()))
        }
        c++
    }
    
    fmt.Println("%v : %v digits", n, len(n.String()))
}