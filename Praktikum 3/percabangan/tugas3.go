package main

import "fmt"

func main() {
    var angka int
    fmt.Print("Bilangan: ")
    fmt.Scan(&angka)

    // Mencari faktor
    factors := []int{}
    for i := 1; i <= angka; i++ {
        if angka%i == 0 {
            factors = append(factors, i)
        }
    }
    fmt.Println("Faktor:", factors)

    // Mengecek prima
    isPrime := true
    for i := 2; i*i <= angka; i++ {
        if angka%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Printf("Prima: %t\n", isPrime)
}