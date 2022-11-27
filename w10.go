package main

import "fmt"

func main() {
    no1()
    // no2()
    // no3()
    // no4()
    // no5()
    // no6()
    // no7()
    // no8()
    // no9()
    // no10()
}

func no1() {
    var x, vokal string
    var valid bool

    fmt.Scan(&x)

    vokal = "AIUEOaiueo"
    valid = true

    for i := 0; i <len(vokal); i++ {
        if x[0] == vokal[i] {
            valid = false
        }
    }

    if valid {
        fmt.Println("konsonan")
    } else {
        fmt.Println("bukan konsonan")
    }
}

func no2() {
    var n int

    fmt.Scan(&n)

    if n % 3 == 0 {
        fmt.Println("Kelipatan 3")
    }
    if n % 5 == 0 {
        fmt.Println("Kelipatan 5")
    }
}

func no3() {
    var s1, s2, s3, temp int

    fmt.Scan(&s1, &s2, &s3)
    
    if (s1 > s2) {
        temp = s1
        s1 = s2
        s2 = temp
    }
    if (s2 > s3) {
        temp = s2
        s2 = s3
        s3 = temp
    }
    if (s1 > s2) {
        temp = s1
        s1 = s2
        s2 = temp
    }

    if s1 == s2 && s2 == s3 {
        fmt.Println("Segitiga Sama Sisi")
    } else if s1 == s2 && s2 != s3 {
        fmt.Println("Segitiga Sama Kaki")
    } else {
        fmt.Println("Segitiga Sembarang")
    }
}

func no4() {
    var n int
    fmt.Scan(&n)
    if n < 0 {
        n *= -1
    }
    fmt.Println(n)
}

func no5() {
    var x1, x2, x3, x4, x5 float64
    fmt.Scan(&x1, &x2, &x3, &x4, &x5)

    if x1 < x2 && x2 < x3 && x3 < x4 && x4 < x5 {
        fmt.Println("Stabil naik")
    } else if x1 > x2 && x2 > x3 && x3 > x4 && x4 > x5 {
        fmt.Println("Stabil turun")
    } else {
        fmt.Println("Tidak stabil")
    }
}

func no6() {
    var x1, x2 float64

    fmt.Scan(&x1, &x2)

    if x2 > x1 {
        fmt.Println("Peningkatan sebesar", x2 - x1)
    } else if x1 > x2 {
        fmt.Println("Penuruan sebesar", x1 - x2)
    } else {
        fmt.Println("Tetap")
    }
}

func no7() {
    var n1, n2, n3, n4, temp int
    fmt.Scan(&n1, &n2, &n3, &n4)

    if n1 > n2 {
        temp = n1
        n1 = n2
        n2 = temp
    }
    if n3 > n4 {
        temp = n3
        n3 = n4
        n4 = temp
    }
    if n1 > n3 {
        temp = n1
        n1 = n3
        n3 = temp
    }
    if n2 > n4 {
        temp = n2
        n2 = n4
        n4 = temp
    }
    if n2 > n3 {
        temp = n2
        n2 = n3
        n3 = temp
    }

    fmt.Println(n4, n1)
}

func no8() {
    var hh1, hh2, mm1, mm2 int
    
    fmt.Scan(&hh1, &mm1, &hh2, &mm2)

    if hh1 > hh2 {
        hh2 += 12
    }

    if mm2 < mm1 {
        hh2 -= 1
        mm2 += 60
    }

    fmt.Printf("%d jam %d menit\n", hh2 - hh1, mm2 - mm1)
}

func no9() {
    var n, cost int
    var make_card, discount, get_card, cashback bool

    fmt.Scan(&n, &make_card)

    discount = n >= 100000
    get_card = make_card && discount
    cashback = n >= 200000 && get_card
    cost = n

    if discount {
        cost = (cost * 9) / 10
    }

    if cashback {
        cost -= 75000
    }

    fmt.Println("Kartu?", get_card)
    fmt.Println("Diskon?", discount)
    fmt.Println("Cashback?", cashback)
    fmt.Println("Rp.", cost)
}

func no10() {
    var n int

    fmt.Scan(&n)
    fmt.Printf("Dewasa ")

    if n <= 15 {
        if n % 5 == 0 {
            fmt.Print(n / 5)
            } else {
            fmt.Print(n / 5 + 1)
        }
    } else {
        fmt.Print(3)
    }

    n -= 15

    if n > 0 {
        fmt.Printf(", kecil ")

        if n <= 10 {
            if n % 2 == 0 {
                fmt.Print(n / 2)
            } else {
                fmt.Print(n / 2 + 1)
            }
        } else {
            fmt.Print(5)
        }
    }

    n -= 10

    if n > 0 {
        fmt.Printf(", dan %d tak berangkat", n)
    }

    fmt.Println()
}
