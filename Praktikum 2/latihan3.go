package main

import (
	"fmt"
)

func main() {
	var jariJari float64
	fmt.Print("Jejari = ")
	fmt.Scanln(&jariJari)

	// Menghitung volume
	volume := (4.0 / 3.0) * 3.1415926535 * jariJari * jariJari * jariJari

	// Menghitung luas permukaan
	luas := 4 * 3.1415926535 * jariJari * jariJari

	// Menampilkan hasil
	fmt.Printf("Bola dengan jejari %.4f memiliki volume %.4f dan luas kulit %.4f\n", jariJari, volume, luas)
}