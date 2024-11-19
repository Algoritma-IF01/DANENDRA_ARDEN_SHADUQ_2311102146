package main

import (
	"fmt"
)

func main() {
	//Membaca masukan untuk jumlah ikan dan kapasitas wadah
	var x, y int
	fmt.Print("Masukkan jumlah ikan: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah ikan dalam satu wadah: ")
	fmt.Scan(&y)

	//Validasi jumlah ikan (x) dan kapasitas wadah (y)
	if x > 1000 {
		fmt.Println("Error: Jumlah ikan tidak boleh lebih dari 1000!")
		return
	}
	if y <= 0 {
		fmt.Println("Error: Kapasitas wadah harus lebih besar dari 0!")
		return
	}

	//Membuat array dengan kapasitas 1000
	var weights [1000]float64

	//Membaca berat ikan sebanyak x
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	//Membuat array untuk menyimpan total berat per wadah
	totalContainers := (x + y - 1) / y //Menghitung jumlah wadah
	containerWeights := make([]float64, totalContainers)

	//Menghitung total berat di setiap wadah
	for i := 0; i < x; i++ {
		containerIndex := i / y
		containerWeights[containerIndex] += weights[i]
	}

	//Menampilkan total berat per wadah
	fmt.Println("Berat total setiap wadah:")
	for i, weight := range containerWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, weight)
	}

	//Menghitung rata-rata berat per wadah
	totalWeight := 0.0
	for _, weight := range containerWeights {
		totalWeight += weight
	}
	averageWeight := totalWeight / float64(totalContainers)

	//Menampilkan rata-rata berat
	fmt.Printf("Berat rata-rata per wadah: %.2f\n", averageWeight)
}
