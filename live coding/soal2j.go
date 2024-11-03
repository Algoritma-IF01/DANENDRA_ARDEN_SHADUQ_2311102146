package main

import (
	"fmt"
)

func hitungBiaya(jumlahMenu int) int {
	const (
		biayaDasar      = 10000
		biayaTambahan   = 2500
		batasMenu       = 50
		biayaMax        = 100000
	)

	if jumlahMenu <= 3 {
		return biayaDasar
	} else if jumlahMenu > batasMenu {
		return biayaMax
	} else {
		return biayaDasar + (jumlahMenu-3)*biayaTambahan
	}
}

func main() {
	var M int
	fmt.Println("Masukkan jumlah rombongan:")
	fmt.Scan(&M)

	for i := 0; i < M; i++ {
		var jumlahMenu, banyakOrang int
		var sisaMakanan bool

		fmt.Printf("Masukkan jumlah menu, banyak orang, dan sisa makanan (true/false) untuk rombongan %d:\n", i+1)
		fmt.Scan(&jumlahMenu, &banyakOrang, &sisaMakanan)

		totalBiaya := hitungBiaya(jumlahMenu)
		if sisaMakanan {
			totalBiaya *= banyakOrang
		}

		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i+1, totalBiaya)
	}
}

