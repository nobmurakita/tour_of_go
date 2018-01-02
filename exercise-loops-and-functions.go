package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := x
    for {
        zz := z - (z * z - x) / (2 * z)
        if math.Abs(zz - z) < 0.00001 {
            break
        }
        z = zz
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
