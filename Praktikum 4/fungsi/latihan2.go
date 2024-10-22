package main

import "fmt"

// Fungsi f(x) = x^2
func f(x int) int {
    return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
    return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
    return x + 1
}

// Fungsi komposisi f(g(h(x)))
func fogoh(x int) int {
    return f(g(h(x)))
}

// Fungsi komposisi g(h(f(x)))
func gohof(x int) int {
    return g(h(f(x)))
}

// Fungsi komposisi h(f(g(x)))
func hofog(x int) int {
    return h(f(g(x)))
}

// Fungsi utama
func main() {
    var a, b, c int

    // Input 3 nilai integer a, b, c
    fmt.Scanf("%d %d %d", &a, &b, &c)

    // Hitung dan cetak hasil komposisi
    fmt.Println("fogoh =", fogoh(a))  // f(g(h(a)))
    fmt.Println("gohof =",gohof(b))  // g(h(f(b)))
    fmt.Println("hofog =",hofog(c))  // h(f(g(c)))
}
