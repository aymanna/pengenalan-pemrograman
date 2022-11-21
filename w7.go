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
    var t, jt, mt, dt int

    fmt.Scan(&t)

    jt = (t / 3600) % 60
    mt = (t / 60) % 60
    dt = t % 60

    fmt.Printf("%d jam %d menit dan %d detik\n", jt, mt, dt)
}

func no2() {
    var x, d1, d3, d4, d23 int
    var discount, cashback bool

    fmt.Scan(&x)    // x = 4567
    d1 = (x / 1000) % 10
    d3 = (x / 10) % 10
    d4 = x % 10
    d23 = (x / 10) % 100

    discount = d23 % 2 == 0

    if d4 == 0 {
        cashback = false
    } else {
        cashback = (d1 + d3) % d4 == 0
    }

    fmt.Println("Diskon?", discount)
    fmt.Println("Cashback?", cashback)
}

func no3() {
    var x, total int

    fmt.Scan(&x)
    total = 0

    for x > 0 {
        total += x % 10
        x /= 10
    }

    fmt.Println(total)
}

func no4() {
    var x, a, b int
    var valid bool

    fmt.Scan(&x)
    valid = true

    for x / 10 != 0 && valid {
        a = x % 10
        x /= 10
        b = x % 10
        valid = a - b == 1 || a - b == -1
    }

    fmt.Println(valid)
}

func no5() {
    var n, a, b, c int

    fmt.Scan(&n)
    a = 0
    b = 1

    for n >= 0 {
        fmt.Printf("%d ", a)
        c = a + b
        a = b
        b = c
        n--
    }
    fmt.Println()
}
