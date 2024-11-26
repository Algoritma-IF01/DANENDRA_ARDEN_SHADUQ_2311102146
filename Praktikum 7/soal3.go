package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan insertion sort
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke posisi berikutnya
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// Fungsi untuk memeriksa jarak antara elemen-elemen array
func checkDistance(arr []int) string {
	if len(arr) <= 1 {
		return "Data berjarak tidak tetap"
	}

	distance := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != distance {
			return "Data berjarak tidak tetap"
		}
	}
	return fmt.Sprintf("Data berjarak %d", distance)
}

func main() {
	var input int
	var data []int

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	if len(data) == 0 {
		fmt.Println("Tidak ada data untuk diproses.")
		return
	}

	// Urutkan array menggunakan insertion sort
	sortedData := insertionSort(data)

	// Periksa jarak antara elemen
	status := checkDistance(sortedData)

	// Cetak keluaran
	fmt.Println("Keluaran:")
	for _, v := range sortedData {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	fmt.Println(status)
}
