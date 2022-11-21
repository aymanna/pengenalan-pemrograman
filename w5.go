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
    var n int
    var text string
    fmt.Scan(&n, &text)

    for i := 0; i < n; i++ {
        fmt.Println(text)
    }
}

func no2() {
    var n, s int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Scan(&s)
        fmt.Println(s * s, 4 * s)
    }
}

func no3() {
    var n, x, sum float64
    sum = 0
    fmt.Scan(&n)

    for i := .0; i < n; i++ {
        fmt.Scan(&x)
        sum += x
    }
    fmt.Printf("%.2f\n", sum / n)
}

func no4() {
    var n, ans int
    ans = 1
    fmt.Scan(&n)

    for i := n; i > 0; i-- {
        ans *= i
    }

    fmt.Println(ans)
}

func no5() {
    var n int
    fmt.Scan(&n)

    for i := 1; i <= n; i++ {
        fmt.Println(i, n % i == 0)
    }
}
