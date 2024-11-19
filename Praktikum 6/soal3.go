package main

import (
	"fmt"
)

// Definisi array balita
type arrBalita [100]float64

// Fungsi untuk mencari nilai minimum dan maksimum dalam array
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arrBerat arrBalita
	var bMin, bMax float64

	// Input jumlah balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n > 100 {
		fmt.Println("Error: Jumlah balita harus antara 1 hingga 100.")
		return
	}

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	hitungMinMax(arrBerat, n, &bMin, &bMax)
	rerata := hitungRerata(arrBerat, n)

	// Tampilkan hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
