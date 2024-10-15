package main

import (
	"fmt"
)

func main() {
	const biayaPerKg = 10000
	const biayaTambahanDiAtas500Gram = 5
	const biayaTambahanDiBawah500Gram = 15

	var beratDalamGram int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratDalamGram)

	// Konversi berat ke kg dan gram
	kg := beratDalamGram / 1000
	gram := beratDalamGram % 1000

	// Hitung biaya dasar
	biayaDasar := kg * biayaPerKg

	// Hitung biaya tambahan
	var biayaTambahan int
	if beratDalamGram > 10000 {
		biayaTambahan = 0
	} else if gram >= 500 {
		biayaTambahan = gram * biayaTambahanDiAtas500Gram
	} else {
		biayaTambahan = gram * biayaTambahanDiBawah500Gram
	}

	// Hitung total biaya
	totalBiaya := biayaDasar + biayaTambahan

	// Buat detail output
	detailBerat := fmt.Sprintf("%d kg + %d gr", kg, gram)
	detailBiaya := fmt.Sprintf("Rp. %d + Rp. %d", biayaDasar, biayaTambahan)

	fmt.Printf("Berat parsel (gram): %d\nDetail berat: %s\nDetail biaya: %s\nTotal biaya: Rp. %d\n", beratDalamGram, detailBerat, detailBiaya, totalBiaya)
}
