package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menampilkan seluruh isi array
func tampilkanArray(array []int) {
	fmt.Println("a. Isi array:", array)
}

// Fungsi untuk menampilkan elemen dengan indeks ganjil
func tampilkanIndeksGanjil(array []int, N int) {
	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks genap
func tampilkanIndeksGenap(array []int, N int) {
	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks kelipatan X
func tampilkanKelipatanX(array []int, N int, X int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", X)
	for i := X; i < N; i += X {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen pada indeks tertentu
func hapusElemen(array []int, N int, indexToDelete int) ([]int, int) {
	if indexToDelete >= 0 && indexToDelete < N {
		array = append(array[:indexToDelete], array[indexToDelete+1:]...)
		N-- // Mengurangi jumlah elemen setelah penghapusan
		fmt.Println("Isi array setelah penghapusan:", array)
	} else {
		fmt.Println("Indeks tidak valid")
	}
	return array, N
}

// Fungsi untuk menghitung rata-rata elemen dalam array
func hitungRataRata(array []int, N int) float64 {
	sum := 0
	for i := 0; i < N; i++ {
		sum += array[i]
	}
	return float64(sum) / float64(N)
}

// Fungsi untuk menghitung standar deviasi elemen dalam array
func hitungStandarDeviasi(array []int, N int) float64 {
	average := hitungRataRata(array, N)
	var varianceSum float64
	for i := 0; i < N; i++ {
		varianceSum += math.Pow(float64(array[i])-average, 2)
	}
	return math.Sqrt(varianceSum / float64(N))
}

// Fungsi untuk menghitung frekuensi kemunculan bilangan tertentu dalam array
func hitungFrekuensi(array []int, N int, searchNum int) int {
	frequency := 0
	for i := 0; i < N; i++ {
		if array[i] == searchNum {
			frequency++
		}
	}
	return frequency
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&N)

	// Membuat array dengan ukuran N dan mengisi nilainya
	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	// a. Menampilkan keseluruhan isi dari array
	tampilkanArray(array)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil
	tampilkanIndeksGanjil(array, N)

	// c. Menampilkan elemen-elemen array dengan indeks genap
	tampilkanIndeksGenap(array, N)

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan X
	var X int
	fmt.Print("d. Masukkan nilai kelipatan (X): ")
	fmt.Scan(&X)
	tampilkanKelipatanX(array, N, X)

	// e. Menghapus elemen pada indeks tertentu
	var indexToDelete int
	fmt.Print("e. Masukkan indeks untuk menghapus elemen: ")
	fmt.Scan(&indexToDelete)
	array, N = hapusElemen(array, N, indexToDelete)

	// f. Menghitung rata-rata dari bilangan di dalam array
	average := hitungRataRata(array, N)
	fmt.Printf("f. Rata-rata: %.2f\n", average)

	// g. Menghitung standar deviasi dari bilangan di dalam array
	stdDev := hitungStandarDeviasi(array, N)
	fmt.Printf("g. Standar deviasi: %.2f\n", stdDev)

	// h. Menghitung frekuensi dari suatu bilangan tertentu
	var searchNum int
	fmt.Print("h. Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&searchNum)
	frequency := hitungFrekuensi(array, N, searchNum)
	fmt.Printf("Frekuensi bilangan %d di dalam array: %d\n", searchNum, frequency)
}
