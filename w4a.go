package main

import "fmt"

func main() {
    var x, y rune
    var are_alphabets, is_same bool

    fmt.Scanf("%c %c\n", &x, &y)

    are_alphabets = is_alphabet(x) && is_alphabet(y)
    is_same = x - y == 32 || y - x == 32

    fmt.Println(are_alphabets && is_same)
}

func is_alphabet(char rune) bool {
    var is_uppercase, is_lowercase bool

    is_uppercase = char >= 'A' && char <= 'Z'
    is_lowercase = char >= 'a' && char <= 'z'

    return is_uppercase || is_lowercase
}
