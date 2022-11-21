package main

import "fmt"

func main() {
    no1()
    // no2()
    // no3()
    // no4()
    // no5()
    // no6()
}

func no1() {
    var username, password string
    var n int

    n = -1

    for !(username == "admin" && password == "admin") {
        fmt.Scan(&username, &password)
        n++
    }

    fmt.Println(n, "\nLogin berhasil")
}

func no2() {
    var balance, transaction int

    transaction = 1     // arbitrary integer other than 0
    fmt.Scan(&balance)

    for transaction != 0 {
        fmt.Scan(&transaction)
        balance += transaction
    }

    fmt.Println(balance)
}

func no3() {
    var x, total int

    total = 0
    fmt.Scan(&x)

    for x != 0 {
        fmt.Printf("%d ", x % 10)
        total += x % 10
        x /= 10
    }

    fmt.Println()
    fmt.Println(total)
}

func no4() {
    var n, m, x, y, total int

    total = 0
    fmt.Scan(&n, &m, &x, &y)

    for n >= x && m >= y {
        n -= x
        m -= y
        total++
    }

    fmt.Println(total)
}

func no5() {
    var x, a, b int
    var valid bool

    valid = true
    fmt.Scan(&x)

    for x / 10 != 0 && valid {
        a = x % 10
        x /= 10
        b = x % 10
        valid = a - b == 1 || a - b == -1
    }

    fmt.Println(valid)
}

func no6() {
    var t, v int

    fmt.Scan(&t)

    for t > 0 {
        fmt.Scan(&v)
        t -= v
        fmt.Println(t <= 0)
    }
}
