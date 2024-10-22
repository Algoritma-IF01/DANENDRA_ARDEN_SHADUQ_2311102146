package main

import "fmt"

// Prosedur untuk mencetak deret berdasarkan aturan
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n) // Cetak suku saat ini
		if n%2 == 0 {        // Jika genap, bagi 2
			n = n / 2
		} else {             // Jika ganjil, gunakan rumus 3n + 1
			n = 3*n + 1
		}
	}
	fmt.Print("1") // Cetak suku terakhir (selalu berakhir di 1)
}

func main() {
	var n int

	// Membaca input suku awal
	fmt.Scan(&n)

	// Memastikan input valid (positif dan kurang dari 1 juta)
	if n > 0 && n < 1000000 {
		cetakDeret(n) // Panggil prosedur cetakDeret
	} else {
		fmt.Println("Masukkan harus bilangan positif kurang dari 1000000.")
	}
}
