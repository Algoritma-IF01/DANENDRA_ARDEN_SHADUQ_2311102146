package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan selection sort
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func main() {
	var n int
	fmt.Println("Masukkan jumlah daerah kerabat:")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("Jumlah daerah harus di antara 1 dan 999.")
		return
	}

	// Array untuk menyimpan nomor rumah setiap daerah
	var daerahRumah [][]int

	// Input data rumah untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah dan nomor rumah untuk daerah %d:\n", i+1)
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("Jumlah rumah kerabat harus di antara 1 dan 999999.")
			return
		}

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		// Simpan data ke array
		daerahRumah = append(daerahRumah, rumah)
	}

	// Proses dan tampilkan hasil setelah semua input selesai
	fmt.Println("\nKeluaran:")
	for _, rumah := range daerahRumah {
		// Urutkan menggunakan selection sort
		sortedRumah := selectionSort(rumah)
		for _, num := range sortedRumah {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}