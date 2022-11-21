package main

import "fmt"

func main() {
    no1()
    // no2()
    // no3()
    // no4()
}

func no1() {
    var r float64
    fmt.Scanln(&r)
    fmt.Println(4 * 22.0/7.0 * r * r)
}

func no2() {
    var x int
    fmt.Scanln(&x)
    d1 := x / 10
    d2 := x % 10
    fmt.Printf("%d%d%d%d\n", d1, d1, d2, d2)
}

func no3() {
    var x string
    fmt.Scanln(&x)
    char := int(x[0])
    fmt.Println(char >= 65 && char <= 90)
}

func no4() {
    var e0, c, e1 int
    fmt.Scanln(&e0, &c, &e1)
    fmt.Println((e0 - e1) / c)
}
