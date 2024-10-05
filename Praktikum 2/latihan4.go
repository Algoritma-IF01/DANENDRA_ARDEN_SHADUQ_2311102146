package main

import "fmt"

func main() {
    var celsius float64

    fmt.Print("Masukkan suhu dalam derajat Celsius: ")
    fmt.Scanln(&celsius)

    // Konversi ke Fahrenheit
    fahrenheit := (celsius * 9 / 5) + 32

    // Konversi ke Reamur
    reamur := celsius * 4 / 5

    // Konversi ke Kelvin
    kelvin := (fahrenheit + 459.67) * 5 / 9

    fmt.Printf("Suhu dalam derajat Fahrenheit: %.2f\n", fahrenheit)
    fmt.Printf("Suhu dalam derajat Reamur: %.2f\n", reamur)
    fmt.Printf("Suhu dalam derajat Kelvin: %.2f\n", kelvin)
}