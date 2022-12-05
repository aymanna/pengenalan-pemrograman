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
}

func no1() {
    var age int
    var kk bool

    fmt.Scan(&age, &kk)

    if age >= 17 && kk {
        fmt.Println("bisa membuat KTP")
    } else {
        fmt.Println("belum bisa membuat KTP")
    }
}

func no2() {
    var weight, kg, gr, cost int

    fmt.Scan(&weight)

    kg = weight / 1000
    gr = weight % 1000
    cost = kg * 10000

    if kg < 10 {
        if gr < 500 {
            cost += gr * 15
        } else {
            cost += gr * 5
        }
    }

    fmt.Println(cost)
}

func no3() {
    var n1, n2, n3, temp int

    fmt.Scan(&n1, &n2, &n3)

    if n1 > n2 {
        temp = n1
        n1 = n2
        n2 = temp
    }
    if n2 > n3 {
        temp = n2
        n2 = n3
        n3 = temp
    }
    if n1 > n2 {
        temp = n1
        n1 = n2
        n2 = temp
    }

    fmt.Println(n1, n2, n3)
}

func no4() {
    var s string
    var is_fired bool

    is_fired = true

    for i := 0; i < 5; i++ {
        fmt.Scan(&s)

        if s != "kalah" {
            is_fired = false
        }
    }

    if is_fired {
        fmt.Println("dipecat")
    } else {
        fmt.Println("tidak dipecat")
    }
}

func no5() {
    var jabatan string
    var masa_kerja, anak_buah, total int

    fmt.Scan(&jabatan, &masa_kerja, &anak_buah)

    if jabatan == "Staf" {
        if masa_kerja < 5 {
            total = 4000
        } else if masa_kerja <= 10 {
            total = 4000 + anak_buah * 100
        } else {
            total = 5000 + anak_buah * 100
        }
    } else if jabatan == "Manajer" || jabatan == "Manager" {
        if masa_kerja <= 10 {
            total = 8500 + anak_buah * 300
        } else {
            total = 10000 + anak_buah * 300
        }
    } else if jabatan == "Direktur" {
        total = 20000 + anak_buah * 500
    }

    fmt.Println(total)
}

func no6() {
    var gigi int
    var kopling, gas bool

    fmt.Scan(&gigi, &kopling, &gas)

    if !kopling && !gas {
        fmt.Println("Mesin mati")
    } else if kopling {
        if (gigi == 0 && gas) || (gigi != 0 && !gas) {
            fmt.Println("Mesin menyala dan motor tidak berjalan")
        } else {
            fmt.Println("Mesin mati")
        }
    } else {
        if gigi == 0 {
            fmt.Println("Mesin menyala dan motor tidak berjalan")
        } else {
            fmt.Println("Mesin menyala")
        }
    }
}

func no7() {
    var n, x, total int

    fmt.Scan(&n)
    total = 0

    for n > 0 {
        fmt.Scan(&x)

        total += x
        n--
    }
    fmt.Println(total)
}

func no8() {
    var member string
    var x, first_three, last_three int
    var cashback, discount float64
    var all_odd, all_even bool

    fmt.Scan(&member)
    cashback = 0
    discount = 0
    first_three = 0
    last_three = 0
    all_odd = true
    all_even = true

    for i := 0; i < 5; i++ {
        fmt.Scan(&x)

        if x % 2 == 0 {
            all_odd = false
        } else {
            all_even = false
        }

        if i >= 0 && i <= 2 {
            first_three += x
        }

        if i >= 2 {
            last_three += x
        }
    }

    if all_even {
        cashback += 3.1 * float64(first_three)
    } else if all_odd {
        discount += 1.7 * float64(last_three)
    } else {            
        cashback += 3.1 * float64(first_three)
        discount += 1.7 * float64(last_three)
    }

    if member == "Yes" {
        cashback *= 1.15
        discount *= 1.15
    }

    if cashback > 35 {
        cashback = 35
    }

    if discount > 50 {
        discount = 50
    }

    fmt.Printf("Cashback: %.2f%% Diskon %.2f%%\n", cashback, discount)
}

func no9() {
    var hh, mm int
    var km, cost_per_km float64
    var rush_hour, normal_hour bool

    fmt.Scan(&hh, &mm, &km)

    rush_hour = ((hh >= 5) && ((hh == 9 && mm == 0) || (hh <= 8))) || 
                ((hh >= 16) && ((hh == 19 && mm == 0) || (hh <= 18)))
    normal_hour = ((hh >= 9) && (hh <= 15)) ||
                  ((hh >= 19) && ((hh == 22 && mm == 0) || (hh <= 21)))

    if rush_hour {
        if km > 0 && km <= 10 {
            cost_per_km = 5000
        } else if km <= 20 {
            cost_per_km = 4500
        } else {
            cost_per_km = 0
        }
    } else if normal_hour {
        cost_per_km = 4000
    } else {
        cost_per_km = 0
    }

    if cost_per_km == 0 {
        fmt.Println("Maaf, kami tidak bisa melayani pesanan anda.")
    } else {
        fmt.Println(cost_per_km * km)
    }
}
