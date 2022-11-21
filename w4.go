package main

import "fmt"

func main() {
    no1()
    // no2()
    // no3()
    // no4()
    // no5()
}

func no1() {
    var x, y int
    fmt.Scan(&x, &y)
    fmt.Println(y % x == 0)
}

func no2() {
    var a, b, c, d, e int
    fmt.Scan(&a, &b, &c, &d, &e)
    fmt.Println(a + b + c - (d + e))
}

func no3() {
    var x, y int
    fmt.Scan(&x, &y)
    fmt.Println(x * x * x == y)
}

func no4() {
    var m, val float64
    scale := []float64{0.38, 0.91, 1.0, 0.38, 2.37, 0.92, 0.89, 1.13}

    fmt.Scan(&m)

    for i := 0; i < len(scale); i++ {
        val = m * 9.8 * scale[i]
        fmt.Printf("%.2f ", val)
    }

    fmt.Println()
}

func no5() {
    var x1, x2, x3, x4, x5 int
    fmt.Scan(&x1, &x2)

    x3 = x1 + x2
    x4 = x2 + x3
    x5 = x3 + x4

    fmt.Println(x3, x4, x5)
}
