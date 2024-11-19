# <h1 align="center">PRAKTIKUM - 6</h1>
<p align="center">DANENDRA ARDEN SHADUQ - 2311102146</p>

### 1. Latihan 1
```GO
package main

import (
	"fmt"
)

type arrInt [2023]int

// Fungsi untuk mencari indeks dari nilai terkecil
func terkecil_2(tabInt arrInt, n int) int {
	var idx int = 0 // indeks data pertama
	var j int = 1   // pencarian dimulai dari data kedua
	for j < n {
		if tabInt[idx] > tabInt[j] { // cek apakah tabInt[j] lebih kecil dari tabInt[idx]
			idx = j // update idx ke indeks baru yang nilainya lebih kecil
		}
		j = j + 1
	}
	return idx // mengembalikan indeks dari nilai terkecil
}

func main() {
	var n int
	var data arrInt

	// Input jumlah elemen N
	fmt.Print("Masukkan jumlah elemen (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi N agar tidak melebihi kapasitas array
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah elemen harus antara 1 dan 2023")
		return
	}

	// Input elemen-elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	// Panggil fungsi untuk mencari indeks nilai terkecil
	idxTerkecil := terkecil_2(data, n)
	fmt.Printf("Indeks nilai terkecil: %d\n", idxTerkecil)
	fmt.Printf("Nilai terkecil: %d\n", data[idxTerkecil])
}
```

Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/l1.png)

### 2. Latihan 2
```GO
package main

import (
	"fmt"
)

// Mendefinisikan tipe data mahasiswa
type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

// Mendefinisikan array mahasiswa dengan kapasitas 2023
type arrMhs [2023]mahasiswa

// Fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
func IPK_2(T arrMhs, n int) int {
	// idx menyimpan indeks mahasiswa dengan IPK tertinggi sementara
	var idx int = 0
	var j int = 1
	for j < n {
		if T[idx].ipk < T[j].ipk {
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var n int
	var data arrMhs

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi jumlah mahasiswa
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023")
		return
	}

	// Input data mahasiswa
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&data[i].nama)
		fmt.Print("NIM: ")
		fmt.Scan(&data[i].nim)
		fmt.Print("Kelas: ")
		fmt.Scan(&data[i].kelas)
		fmt.Print("Jurusan: ")
		fmt.Scan(&data[i].jurusan)
		fmt.Print("IPK: ")
		fmt.Scan(&data[i].ipk)
	}

	// Panggil fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
	idxTertinggi := IPK_2(data, n)

	// Tampilkan data mahasiswa dengan IPK tertinggi
	fmt.Println("\nMahasiswa dengan IPK tertinggi:")
	fmt.Printf("Nama    : %s\n", data[idxTertinggi].nama)
	fmt.Printf("NIM     : %s\n", data[idxTertinggi].nim)
	fmt.Printf("Kelas   : %s\n", data[idxTertinggi].kelas)
	fmt.Printf("Jurusan : %s\n", data[idxTertinggi].jurusan)
	fmt.Printf("IPK     : %.2f\n", data[idxTertinggi].ipk)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/l2.png)

### 3. Soal 1
```GO
package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Print("Masukkan jumlah anak kelinci: ")
    fmt.Scan(&n)

    if n <= 0 || n > 1000 {
        fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000")
        return
    }

    weights := make([]float64, n)
    fmt.Println("Masukkan berat anak kelinci:")
    for i := 0; i < n; i++ {
        fmt.Scan(&weights[i])
    }

    minWeight, maxWeight := weights[0], weights[0]

    for _, weight := range weights[1:] {
        if weight < minWeight {
            minWeight = weight
        }
        if weight > maxWeight {
            maxWeight = weight
        }
    }

    fmt.Printf("Berat kelinci terkecil: %.2f\n", minWeight)
    fmt.Printf("Berat kelinci terbesar: %.2f\n", maxWeight)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/s1.png)

### 4. Soal 2
```GO
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

```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/s2.png)


### 5. Soal 3
```GO
package main

import (
	"fmt"
)

// Definisi array balita
type arrBalita [100]float64

// Fungsi untuk mencari nilai minimum dan maksimum dalam array
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var arrBerat arrBalita
	var bMin, bMax float64

	// Input jumlah balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi input
	if n <= 0 || n > 100 {
		fmt.Println("Error: Jumlah balita harus antara 1 hingga 100.")
		return
	}

	// Input berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}

	// Hitung nilai minimum, maksimum, dan rata-rata
	hitungMinMax(arrBerat, n, &bMin, &bMax)
	rerata := hitungRerata(arrBerat, n)

	// Tampilkan hasil
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/s3.png)
