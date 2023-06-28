package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    var num int
    fmt.Print("Digite um número inteiro: ")
    fmt.Scan(&num)
    fmt.Printf("O fatorial de %d é %d\n", num, factorial(num))
}
