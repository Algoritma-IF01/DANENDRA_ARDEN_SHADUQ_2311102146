package main

import (
	"fmt"
)

func sumMultiplesOfFour(total int, input int) int {
	if input < 0 {
		return total
	}
	if input > 0 && input%4 == 0 {
		total += input
	}
	var nextInput int
	fmt.Scan(&nextInput)
	return sumMultiplesOfFour(total, nextInput)
}

func main() {
	var input int

	fmt.Println("Masukkan bilangan bulat (bilangan negatif untuk berhenti):")
	fmt.Scan(&input)

	result := sumMultiplesOfFour(0, input)
	fmt.Printf("Jumlah bilangan bulat positif kelipatan 4: %d\n", result)
}