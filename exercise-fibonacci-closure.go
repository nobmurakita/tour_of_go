package main

import "fmt"

// fibonacci is a function that returns
// a function that returns as int.
func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        n := a
        a, b = b, a + b
        return n
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
