package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

// Fungsi untuk memeriksa apakah suatu titik berada di dalam lingkaran
func dalamLingkaran(cx, cy, r, x, y int) bool {
	return jarak(cx, cy, x, y) <= float64(r)
}

func main() {
	// Input lingkaran 1
	var cx1, cy1, r1 int
	fmt.Scan(&cx1, &cy1, &r1)

	// Input lingkaran 2
	var cx2, cy2, r2 int
	fmt.Scan(&cx2, &cy2, &r2)

	// Input titik sembarang
	var x, y int
	fmt.Scan(&x, &y)

	// Memeriksa posisi titik terhadap kedua lingkaran
	diDalam1 := dalamLingkaran(cx1, cy1, r1, x, y)
	diDalam2 := dalamLingkaran(cx2, cy2, r2, x, y)

	// Menentukan keluaran berdasarkan kondisi
	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
