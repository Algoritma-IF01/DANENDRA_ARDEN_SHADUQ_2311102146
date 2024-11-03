package main

import (
	"fmt"
)

func main() {
	// Meminta input nama klub yang bertanding
	var klubA, klubB string
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klubB)

	// Array untuk menyimpan hasil pemenang setiap pertandingan
	var hasilPertandingan []string

	// Variabel untuk menyimpan skor
	var skorA, skorB int
	pertandingan := 1

	// Loop untuk memasukkan skor setiap pertandingan
	for {
		// Menerima input skor dengan format "skorA skorB" dalam satu baris
		fmt.Printf("Pertandingan %d - Masukkan skor (%s %s): ", pertandingan, klubA, klubB)
		_, err := fmt.Scanln(&skorA, &skorB)
		if err != nil {
			fmt.Println("Input tidak valid, program dihentikan.")
			break
		}

		// Cek apakah skor valid (tidak negatif)
		if skorA < 0 || skorB < 0 {
			fmt.Println("Skor tidak valid, program dihentikan.")
			break
		}

		// Menentukan pemenang dan menyimpan hasilnya ke array hasilPertandingan
		if skorA > skorB {
			hasilPertandingan = append(hasilPertandingan, fmt.Sprintf("Hasil %d: %s", pertandingan, klubA))
		} else if skorB > skorA {
			hasilPertandingan = append(hasilPertandingan, fmt.Sprintf("Hasil %d: %s", pertandingan, klubB))
		} else {
			hasilPertandingan = append(hasilPertandingan, fmt.Sprintf("Hasil %d: Draw", pertandingan))
		}

		// Naikkan nomor pertandingan
		pertandingan++
	}

	// Menampilkan daftar hasil pertandingan setelah input selesai
	fmt.Println("Daftar hasil pertandingan:")
	for _, hasil := range hasilPertandingan {
		fmt.Println(hasil)
	}
	fmt.Println("Pertandingan selesai")
}
