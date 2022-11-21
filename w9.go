package main

import "fmt"

func main() {
    no1()
    // no2()
    // no3()
    // no4()
    // no5()
    // no6()
    // no10()
}

func no1() {
    var x byte
    var ascii_value int
    var is_numerical, is_lowercase, is_uppercase, valid bool

    fmt.Scanf("%c\n", &x)
    ascii_value = int(x)

    is_numerical = ascii_value >= '0' && ascii_value <= '9'
    is_lowercase = ascii_value >= 'a' && ascii_value <= 'z'
    is_uppercase = ascii_value >= 'A' && ascii_value <= 'Z'

    valid = is_numerical || is_lowercase || is_uppercase
    fmt.Println(valid)
}

func no2() {
    var n int
    fmt.Scan(&n)
    fmt.Println((n % 4 == 0 && n % 100 != 0) || (n % 400 == 0))
}

func no3() {
    var s1, s2, s3, temp int
    var valid bool

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

    valid = s1 + s2 > s3
    fmt.Println(valid)
}

func no4() {
    var n int
    var make_card, card, discount, cashback bool

    fmt.Scan(&n, &make_card)

    discount = n >= 100000
    card = make_card && discount
    cashback = n >= 200000 && card

    fmt.Println("Kartu?", card)
    fmt.Println("Diskon?", discount)
    fmt.Println("Cashback?", cashback)
}

func no5() {
    var n1, n2, n3, temp int
    var valid bool

    fmt.Scan(&n1, &n2, &n3)

    if (n1 > n2) {
        temp = n1
        n1 = n2
        n2 = temp
    }
    if (n2 > n3) {
        temp = n2
        n2 = n3
        n3 = temp
    }
    if (n1 > n2) {
        temp = n1
        n1 = n2
        n2 = temp
    }

    valid = (n1 + n3) / 2 == n2 && (n1 + n3) % 2 != 1
    fmt.Println(valid)
}

func no6() {
    var r1, r2, d int
    fmt.Scan(&r1, &r2, &d)
    fmt.Println(r1 + r2 <= d)
}

func no10() {
    var n int
    var x1, x2, x3, x4, x5, valid bool

    valid = true
    fmt.Scan(&n)

    for n > 0 {
        n--
        fmt.Scan(&x1, &x2, &x3, &x4, &x5)

        if !(x1 && x2 && x3 && x4 && x5) {
            valid = false
        }
    }

    fmt.Println(valid)
}
