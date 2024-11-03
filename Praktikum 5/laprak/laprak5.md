# <h1 align="center">PRAKTIKUM 5</h1>
<p align="center">DANENDRA ARDEN SHADUQ - 2311102146</p>

### 1. Latihan 1
```GO
package main

import "fmt"

type bilangan int
type pecahan float64

func main() {
	var a, b bilangan
	var hasil pecahan
	a = 9
	b = 5
	hasil = pecahan(a) / pecahan(b)
	fmt.Println(hasil)

}
```

Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%205/ss/l1.png)

### 2. Latihan 2
```GO
package main
import "fmt"
type waktu struct {
	jam, menit, detik int
}

func main(){
	var wParkir, wPulang, durasi waktu
	var dParkir, dPulang, lParkir int
	fmt.Scan(&wParkir.jam, &wParkir.menit, &wParkir.detik)
	fmt.Scan(&wPulang.jam, &wPulang.menit, &wPulang.detik)
	dParkir = wParkir.detik + wParkir.menit*60 + wParkir.jam*3600
	dPulang = wPulang.detik + wPulang.menit*60 + wPulang.jam*3600
	lParkir = dPulang - dParkir
	durasi.jam = lParkir / 3600
	durasi.menit = lParkir % 3600 / 60
	durasi.detik = lParkir % 3600 % 60
	fmt.Printf("Lama parkir : %d jam %d menit %d detik",
		durasi.jam, durasi.menit, durasi.detik)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%205/ss/l2.png)

### 3. Latihan 3
```GO
package main

import "fmt"

// Definisi tipe CircType
type CircType struct {
	Radius float64
}

// Definisi tipe NewType
type NewType struct {
	Name string
}

func main() {
	var (
		// Array arr mempunyai 73 elemen, masing-masing bertipe CircType
		arr [73]CircType

		// Array buf dengan 5 elemen, dengan nilai awal 7,3,5,2,11
		buf = [5]byte{7, 3, 5, 2, 11}

		// Array mhs berisi 2000 elemen NewType
		mhs [2000]NewType

		// Array dua dimensi rec berisi nilai float64
		rec [20][40]float64
	)

	// Mengisi beberapa elemen contoh
	arr[0] = CircType{Radius: 5.5}
	mhs[0] = NewType{Name: "Alice"}
	rec[0][0] = 3.14

	// Contoh penggunaan variabel
	fmt.Println("arr[0]:", arr[0])
	fmt.Println("buf:", buf)
	fmt.Println("mhs[0]:", mhs[0])
	fmt.Println("rec[0][0]:", rec[0][0])
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%205/ss/l3.png)

### 4. Latihan 4
```GO
package main

import "fmt"

func main() {
	// Membuat map dengan tipe string sebagai kunci dan integer sebagai nilai
	scores := map[string]int{
		"John": 90,
		"Anne": 85,
	}

	// Mengambil nilai dari kunci
	johnScore := scores["John"]
	fmt.Println("Nilai John:", johnScore)

	// Mengganti nilai dari kunci yang ada
	scores["John"] = 95
	fmt.Println("Nilai John setelah diganti:", scores["John"])

	// Menambah kunci baru dengan nilai
	scores["David"] = 88
	fmt.Println("Nilai David ditambahkan:", scores["David"])

	// Menghapus kunci dari map
	delete(scores, "Anne")
	fmt.Println("Map setelah kunci 'Anne' dihapus:", scores)

	// Mengecek apakah kunci ada dalam map
	if score, ada := scores["David"]; ada {
		fmt.Println("Nilai David ditemukan:", score)
	} else {
		fmt.Println("Nilai David tidak ditemukan")
	}
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%205/ss/l4.png)


### 5. Soal 1
```GO
package main

import (
	"fmt"
	"math"
)

// Definisi tipe bentukan untuk titik
type Titik struct {
	x int
	y int
}

// Definisi tipe bentukan untuk lingkaran
type Lingkaran struct {
	center Titik
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(p Titik, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func didalam(c Lingkaran, p Titik) bool {
	return jarak(p, c.center) < float64(c.radius)
}

func main() {
	var (
		// Mengambil input untuk lingkaran 1
		lingkaran1 Lingkaran
		// Mengambil input untuk lingkaran 2
		lingkaran2 Lingkaran
		// Mengambil input untuk titik sembarang
		point Titik
	)

	// Input untuk lingkaran 1 (cx, cy, r)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scan(&lingkaran1.center.x, &lingkaran1.center.y, &lingkaran1.radius)

	// Input untuk lingkaran 2 (cx, cy, r)
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scan(&lingkaran2.center.x, &lingkaran2.center.y, &lingkaran2.radius)

	// Input untuk titik sembarang (x, y)
	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&point.x, &point.y)

	// Mengecek posisi titik terhadap kedua lingkaran
	inLingkaran1 := didalam(lingkaran1, point)
	inLingkaran2 := didalam(lingkaran2, point)

	if inLingkaran1 && inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%205/ss/sl1.png)

### 6. Soal 2
```GO
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menampilkan seluruh isi array
func tampilkanArray(array []int) {
	fmt.Println("a. Isi array:", array)
}

// Fungsi untuk menampilkan elemen dengan indeks ganjil
func tampilkanIndeksGanjil(array []int, N int) {
	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks genap
func tampilkanIndeksGenap(array []int, N int) {
	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks kelipatan X
func tampilkanKelipatanX(array []int, N int, X int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", X)
	for i := X; i < N; i += X {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen pada indeks tertentu
func hapusElemen(array []int, N int, indexToDelete int) ([]int, int) {
	if indexToDelete >= 0 && indexToDelete < N {
		array = append(array[:indexToDelete], array[indexToDelete+1:]...)
		N-- // Mengurangi jumlah elemen setelah penghapusan
		fmt.Println("Isi array setelah penghapusan:", array)
	} else {
		fmt.Println("Indeks tidak valid")
	}
	return array, N
}

// Fungsi untuk menghitung rata-rata elemen dalam array
func hitungRataRata(array []int, N int) float64 {
	sum := 0
	for i := 0; i < N; i++ {
		sum += array[i]
	}
	return float64(sum) / float64(N)
}

// Fungsi untuk menghitung standar deviasi elemen dalam array
func hitungStandarDeviasi(array []int, N int) float64 {
	average := hitungRataRata(array, N)
	var varianceSum float64
	for i := 0; i < N; i++ {
		varianceSum += math.Pow(float64(array[i])-average, 2)
	}
	return math.Sqrt(varianceSum / float64(N))
}

// Fungsi untuk menghitung frekuensi kemunculan bilangan tertentu dalam array
func hitungFrekuensi(array []int, N int, searchNum int) int {
	frequency := 0
	for i := 0; i < N; i++ {
		if array[i] == searchNum {
			frequency++
		}
	}
	return frequency
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&N)

	// Membuat array dengan ukuran N dan mengisi nilainya
	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	// a. Menampilkan keseluruhan isi dari array
	tampilkanArray(array)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil
	tampilkanIndeksGanjil(array, N)

	// c. Menampilkan elemen-elemen array dengan indeks genap
	tampilkanIndeksGenap(array, N)

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan X
	var X int
	fmt.Print("d. Masukkan nilai kelipatan (X): ")
	fmt.Scan(&X)
	tampilkanKelipatanX(array, N, X)

	// e. Menghapus elemen pada indeks tertentu
	var indexToDelete int
	fmt.Print("e. Masukkan indeks untuk menghapus elemen: ")
	fmt.Scan(&indexToDelete)
	array, N = hapusElemen(array, N, indexToDelete)

	// f. Menghitung rata-rata dari bilangan di dalam array
	average := hitungRataRata(array, N)
	fmt.Printf("f. Rata-rata: %.2f\n", average)

	// g. Menghitung standar deviasi dari bilangan di dalam array
	stdDev := hitungStandarDeviasi(array, N)
	fmt.Printf("g. Standar deviasi: %.2f\n", stdDev)

	// h. Menghitung frekuensi dari suatu bilangan tertentu
	var searchNum int
	fmt.Print("h. Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&searchNum)
	frequency := hitungFrekuensi(array, N, searchNum)
	fmt.Printf("Frekuensi bilangan %d di dalam array: %d\n", searchNum, frequency)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%205/ss/sl2.png)

### 7. Soal 3
```GO
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
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%205/ss/sl3.png)
### 8. Soal 4
```GO
package main

import "fmt"

const NMAX int = 127
type tabel [NMAX]rune

var tab tabel
var m int

// Fungsi untuk mengisi array dari input pengguna
func isiArray(t *tabel, n *int) {
	// I.S. Data tersedia dalam piranti masukan
	// F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
	// Proses input selama karakter bukanlah TITIK dan n <= NMAX
	var input rune
	*n = 0
	for {
		fmt.Print("Masukkan karakter (akhiri dengan titik '.'): ")
		fmt.Scanf("%c\n", &input)
		if input == '.' {
			break
		}
		t[*n] = input
		*n++
		if *n >= NMAX {
			break
		}
	}
}

// Fungsi untuk mencetak array
func cetakArray(t tabel, n int) {
	// I.S. Terdefinisi array t yang berisi sejumlah n karakter
	// F.S. n karakter dalam array muncul di layar
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan isi array
func balikanArray(t *tabel, n int) {
	// I.S. Terdefinisi array t yang berisi sejumlah n karakter
	// F.S. Urutan isi array t terbalik
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	// Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom,
	// dan false apabila sebaliknya
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	// si array tab dengan memanggil prosedur isiArray
	isiArray(&tab, &m)

	// Cetak array tab
	fmt.Print("Teks: ")
	cetakArray(tab, m)

	// Periksa apakah array adalah palindrom
	fmt.Print("Palindrom: ")
	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// Balikkan isi array tab dengan memanggil balikanArray
	balikanArray(&tab, m)

	// Cetak array tab yang sudah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%205/ss/sl4.png)
