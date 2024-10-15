package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	// Initialize the product for sqrt(2) approximation
	approximation := 1.0

	// Loop through k from 0 to K and compute the product
	for k := 0; k <= K; k++ {
		numerator := (4*float64(k) + 2) * (4*float64(k) + 2) // (4k + 2)^2
		denominator := (4*float64(k) + 1) * (4*float64(k) + 3) // (4k + 1)(4k + 3)
		approximation *= numerator / denominator
	}

	// Output the result, approximating sqrt(2)
	fmt.Printf("Nilai akar 2 = %.10f\n", approximation)
}
