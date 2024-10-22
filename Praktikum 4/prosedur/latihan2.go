package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk menghitung skor dan jumlah soal yang diselesaikan
func hitungSkor(
	s1, s2, s3, s4, s5, s6, s7, s8 int,
) (int, int) {
	jumlahSoal := 0
	totalSkor := 0

	// Cek tiap soal secara manual
	if s1 <= 300 {
		jumlahSoal++
		totalSkor += s1
	}
	if s2 <= 300 {
		jumlahSoal++
		totalSkor += s2
	}
	if s3 <= 300 {
		jumlahSoal++
		totalSkor += s3
	}
	if s4 <= 300 {
		jumlahSoal++
		totalSkor += s4
	}
	if s5 <= 300 {
		jumlahSoal++
		totalSkor += s5
	}
	if s6 <= 300 {
		jumlahSoal++
		totalSkor += s6
	}
	if s7 <= 300 {
		jumlahSoal++
		totalSkor += s7
	}
	if s8 <= 300 {
		jumlahSoal++
		totalSkor += s8
	}

	return jumlahSoal, totalSkor
}

func main() {
	var nama string

	pemenang := ""
	maxSoal := -1
	minSkor := 99999

	// Loop untuk membaca input peserta
	for {
		// Baca nama peserta
		fmt.Scan(&nama)
		if strings.ToLower(nama) == "selesai" {
			break
		}

		// Baca waktu penyelesaian 8 soal
		var s1, s2, s3, s4, s5, s6, s7, s8 int
		fmt.Scan(&s1, &s2, &s3, &s4, &s5, &s6, &s7, &s8)

		// Hitung skor dan jumlah soal yang diselesaikan
		jumlahSoal, totalSkor := hitungSkor(s1, s2, s3, s4, s5, s6, s7, s8)

		// Tentukan pemenang berdasarkan jumlah soal dan total skor
		if jumlahSoal > maxSoal || (jumlahSoal == maxSoal && totalSkor < minSkor) {
			pemenang = nama
			maxSoal = jumlahSoal
			minSkor = totalSkor
		}
	}

	// Cetak hasil pemenang
	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minSkor)
}
