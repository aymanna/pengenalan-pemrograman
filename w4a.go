package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y string
    fmt.Scan(&x, &y)

    areAlphabets := isAlphabet(x) && isAlphabet(y)
    isSame := math.Abs(float64(x[0]) - float64(y[0])) == 32

    fmt.Println(areAlphabets && isSame)
}

func isAlphabet(char string) bool {
    ascii_value := int(char[0])
    isUpperCase := ascii_value > 64 && ascii_value < 91
    isLowerCase := ascii_value > 96 && ascii_value < 123

    return isUpperCase || isLowerCase
}
