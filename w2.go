package main

import (
    "fmt"
    "math"
)

func main() {
    no1()
    // no2()
    // no3()
}

func no1() {
    var x, y float64

    fmt.Scanln(&x, &y)
    total := 1/(3*math.Pow(x, 2.0)+10) + 10*y + 7
    fmt.Println(total)
}

func no2() {
    var x float64

    fmt.Scanln(&x)
    p := math.Pow(x, 3.0) + 3*x
    q := math.Pow(x, 4.0) - 3*math.Pow(x, 2.0) + 4

    fmt.Println(p / q)
}

func no3() {
    var x int

    fmt.Scanln(&x)
    d1 := (x / 100) % 10
    d2 := x / 10 % 10
    d3 := x % 10

    fmt.Println(d1, d2, d3)
}
