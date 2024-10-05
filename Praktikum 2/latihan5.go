package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk menyimpan input
	var a, b, c, d, e int
	var ch1, ch2, ch3 rune

	// Membaca 5 angka integer
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)

	// Membaca 3 karakter
	fmt.Scanf("%c%c%c", &ch1, &ch2, &ch3)

	// Cetak karakter dari integer
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	// Geser karakter input satu posisi ke depan
	fmt.Printf("%c%c%c\n", ch1+1, ch2+1, ch3+1)
}
