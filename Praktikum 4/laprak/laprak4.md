# <h1 align="center">PRAKTIKUM 4</h1>
<p align="center">DANENDRA ARDEN SHADUQ - 2311102146</p>

## FUNGSI
### 1. CONTOH PADA FAKTORIAL & PERMUTASI 
```GO
package main
import "fmt"

func faktorial(n int) int{
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++{
		hasil = hasil * i
	}
	return hasil
}
func permutasi(n, r int) int{
	return faktorial(n) / faktorial(n-r)
}

func main(){
	var a,b int
	fmt.Scan(&a, &b)
	if a >= b {
		fmt.Println(permutasi(a, b))
	}else{
		fmt.Println(permutasi(b,a))
	}
}
```

Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%204/ss/cth.png)

### 2. Latihan 1
```GO
package main
import (
	"fmt"
)

//fungsi untuk menghitung faktorial
func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

//fungsi untuk menghitung permutasi
func permutation(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

//fungsi untuk menghitung kombinasi
func combination(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	//meminta input dari pengguna
	fmt.Println("Masukkan nilai a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	//menghitung permutasi dan kombinasi untuk a terhadap b
	p1 := permutation(a, c)
	c1 := combination(a, c)

	//menghitung permutasi dan kombinasi untuk c terhadap d
	p2 := permutation(b, d)
	c2 := combination(b, d)

	//output
	fmt.Printf("P(%d, %d) = %d\n", a, c, p1)
	fmt.Printf("C(%d, %d) = %d\n", a, c, c1)
	fmt.Printf("P(%d, %d) = %d\n", b, d, p2)
	fmt.Printf("C(%d, %d) = %d\n", b, d, c2)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%204/ss/L1.png)

### 3. Latihan 2
```GO
package main

import "fmt"

// Fungsi f(x) = x^2
func f(x int) int {
    return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
    return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
    return x + 1
}

// Fungsi komposisi f(g(h(x)))
func fogoh(x int) int {
    return f(g(h(x)))
}

// Fungsi komposisi g(h(f(x)))
func gohof(x int) int {
    return g(h(f(x)))
}

// Fungsi komposisi h(f(g(x)))
func hofog(x int) int {
    return h(f(g(x)))
}

// Fungsi utama
func main() {
    var a, b, c int

    // Input 3 nilai integer a, b, c
    fmt.Scanf("%d %d %d", &a, &b, &c)

    // Hitung dan cetak hasil komposisi
    fmt.Println("fogoh =", fogoh(a))  // f(g(h(a)))
    fmt.Println("gohof =",gohof(b))  // g(h(f(b)))
    fmt.Println("hofog =",hofog(c))  // h(f(g(c)))
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%204/ss/L2.png)

### 4. Latihan 3
```GO
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
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%204/ss/L3.png)

## PROSEDUR
### 1. Fibonacci
```GO
package main
import (
	"fmt"
)

// fungsi rekursif untuk menghitung deret fobonaci
func fibonaci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonaci(n-1) + fibonaci(n-2)
	}
}

func main() {
	// menambilkan deret fibonaci hingga suku ke-10
	fmt.Println("Deret fibonaci hingga suku ke-10 : ")
	for i := 0 ; i < 10 ; i++ {
		fmt.Printf("fibonaci (%d) = %d\n", i, fibonaci(i))
	}
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%204/ss/fibo.png)

### 2. Latihan 2
```GO
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
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%204/ss/T2.png)

### 3. Latihan 3
```GO
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
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%204/ss/T3.png)
