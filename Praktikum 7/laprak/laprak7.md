# <h1 align="center">PRAKTIKUM - 7</h1>
<p align="center">DANENDRA ARDEN SHADUQ - 2311102146</p>

### 1. Latihan 1
```GO
package main

import "fmt"

type arrInt [4321]int

func insertionSort1(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara mengecil (descending) dengan INSERTION SORT */
	var temp, i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya yang lebih kecil dari temp
		for j > 0 && temp > T[j-1] {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func main() {
	// Contoh penggunaan
	var T arrInt
	n := 5
	T[0], T[1], T[2], T[3], T[4] = 22, 12, 34, 64, 25

	fmt.Println("Array sebelum diurutkan:", T[:n])
	insertionSort1(&T, n)
	fmt.Println("Array setelah diurutkan secara descending:", T[:n])
}
```

Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%207/ss/l1.png)

### 2. Latihan 2
```GO
package main

import "fmt"

type arrInt [4321]int

func selectionSort1(T *arrInt, n int) {
	/* I.S. terdefinisi array T yang berisi n bilangan bulat
	   F.S. array T terurut secara ascending atau membesar dengan SELECTION SORT */
	for i := 0; i < n-1; i++ {
		// Inisialisasi indeks minimum
		idx_min := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[idx_min] {
				idx_min = j
			}
		}
		// Tukar elemen T[i] dengan T[idx_min] jika perlu
		if idx_min != i {
			T[i], T[idx_min] = T[idx_min], T[i]
		}
	}
}

func main() {
	// Contoh penggunaan
	var T arrInt
	n := 5
	T[0], T[1], T[2], T[3], T[4] = 64, 34, 25, 12, 22

	fmt.Println("Array sebelum diurutkan:", T[:n])
	selectionSort1(&T, n)
	fmt.Println("Array setelah diurutkan:", T[:n])
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%207/ss/l2.png)

### 3. Latihan 3
```GO
package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func selectionSort2(T *arrMhs, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara ascending berdasarkan ipk dengan
	   menggunakan algoritma SELECTION SORT */

	var idx_min int
	var temp mahasiswa

	for i := 0; i < n-1; i++ {
		// Inisialisasi indeks minimum
		idx_min = i

		// Cari elemen dengan IPK terkecil di subarray [i+1, n-1]
		for j := i + 1; j < n; j++ {
			if T[j].ipk < T[idx_min].ipk {
				idx_min = j
			}
		}

		// Tukar elemen di indeks i dengan elemen di idx_min jika perlu
		if idx_min != i {
			temp = T[i]
			T[i] = T[idx_min]
			T[idx_min] = temp
		}
	}
}

func main() {
	// Contoh data mahasiswa
	var T arrMhs
	T[0] = mahasiswa{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[1] = mahasiswa{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[2] = mahasiswa{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[3] = mahasiswa{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	selectionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan IPK:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%207/ss/l3.png)

### 4. Latihan 4
```GO
package main

import "fmt"

type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]mahasiswa

func insertionSort2(T *arrMhs, n int) {
	/* I.S. terdefinisi array T yang berisi n data mahasiswa
	   F.S. array T terurut secara mengecil (descending) berdasarkan nama
	   dengan menggunakan algoritma INSERTION SORT */
	var temp mahasiswa
	var i, j int

	for i = 1; i < n; i++ {
		temp = T[i] // Simpan elemen ke-i
		j = i       // Inisialisasi indeks pembanding

		// Geser elemen-elemen sebelumnya
		for j > 0 && temp.nama > T[j-1].nama {
			T[j] = T[j-1]
			j--
		}

		// Tempatkan temp pada posisi yang sesuai
		T[j] = temp
	}
}

func main() {
	// Contoh data mahasiswa
	var T arrMhs
	T[0] = mahasiswa{"Charlie", "125", "A", "Teknik Informatika", 3.5}
	T[1] = mahasiswa{"Alice", "123", "A", "Teknik Informatika", 3.8}
	T[2] = mahasiswa{"Bob", "124", "B", "Sistem Informasi", 3.2}
	T[3] = mahasiswa{"Diana", "126", "B", "Sistem Informasi", 3.9}
	n := 4

	fmt.Println("Data mahasiswa sebelum diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}

	insertionSort2(&T, n)

	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan nama (descending):")
	for i := 0; i < n; i++ {
		fmt.Printf("%s - %s - %s - %s - %.2f\n", T[i].nama, T[i].nim, T[i].kelas, T[i].jurusan, T[i].ipk)
	}
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%207/ss/l4.png)


### 5. Soal 1
```GO
package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan selection sort
func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func main() {
	var n int
	fmt.Println("Masukkan jumlah daerah kerabat:")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("Jumlah daerah harus di antara 1 dan 999.")
		return
	}

	// Array untuk menyimpan nomor rumah setiap daerah
	var daerahRumah [][]int

	// Input data rumah untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah dan nomor rumah untuk daerah %d:\n", i+1)
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("Jumlah rumah kerabat harus di antara 1 dan 999999.")
			return
		}

		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		// Simpan data ke array
		daerahRumah = append(daerahRumah, rumah)
	}

	// Proses dan tampilkan hasil setelah semua input selesai
	fmt.Println("\nKeluaran:")
	for _, rumah := range daerahRumah {
		// Urutkan menggunakan selection sort
		sortedRumah := selectionSort(rumah)
		for _, num := range sortedRumah {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/sl1.png)

### 6. Soal 2
```GO
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan jumlah daerah:")
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text()) // Jumlah daerah

	for i := 0; i < t; i++ {
		fmt.Printf("Masukkan angka untuk daerah %d (pisahkan dengan spasi):\n", i+1)
		scanner.Scan()
		input := scanner.Text()

		// Parsing input menjadi array bilangan bulat
		nums := parseInput(input)

		// Pisahkan bilangan ganjil dan genap
		odds, evens := separateOddEven(nums)

		// Urutkan ganjil membesar dan genap mengecil
		sort.Ints(odds)
		sort.Sort(sort.Reverse(sort.IntSlice(evens)))

		// Cetak hasil
		fmt.Printf("Hasil daerah %d:\n", i+1)
		printResult(odds, evens)
	}
}

// Fungsi untuk memparsing input menjadi array bilangan bulat
func parseInput(input string) []int {
	parts := strings.Fields(input)
	nums := make([]int, len(parts))
	for i, p := range parts {
		nums[i], _ = strconv.Atoi(p)
	}
	return nums
}

// Fungsi untuk memisahkan bilangan ganjil dan genap
func separateOddEven(nums []int) ([]int, []int) {
	var odds, evens []int
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	return odds, evens
}

// Fungsi untuk mencetak hasil
func printResult(odds []int, evens []int) {
	for _, odd := range odds {
		fmt.Print(odd, " ")
	}
	for _, even := range evens {
		fmt.Print(even, " ")
	}
	fmt.Println()
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/sl2.png)

### 7. Soal 3
```GO
package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan insertion sort
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke posisi berikutnya
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// Fungsi untuk memeriksa jarak antara elemen-elemen array
func checkDistance(arr []int) string {
	if len(arr) <= 1 {
		return "Data berjarak tidak tetap"
	}

	distance := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != distance {
			return "Data berjarak tidak tetap"
		}
	}
	return fmt.Sprintf("Data berjarak %d", distance)
}

func main() {
	var input int
	var data []int

	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	if len(data) == 0 {
		fmt.Println("Tidak ada data untuk diproses.")
		return
	}

	// Urutkan array menggunakan insertion sort
	sortedData := insertionSort(data)

	// Periksa jarak antara elemen
	status := checkDistance(sortedData)

	// Cetak keluaran
	fmt.Println("Keluaran:")
	for _, v := range sortedData {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	fmt.Println(status)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/sl3.png)

### 8. Soal 4
```GO
package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Println("Masukkan jumlah buku:")
	fmt.Scanf("%d\n", n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scanf("%d %s %s %s %d %d %d\n",
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	// Menampilkan buku dengan rating tertinggi
	maxRating := -1
	var bukuTerfavorit Buku
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			bukuTerfavorit = pustaka[i]
		}
	}
	fmt.Println("Buku terfavorit:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		bukuTerfavorit.id, bukuTerfavorit.judul, bukuTerfavorit.penulis, bukuTerfavorit.penerbit,
		bukuTerfavorit.tahun, bukuTerfavorit.rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	// Sorting menggunakan Insertion Sort berdasarkan rating
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	// Menampilkan 5 buku dengan rating tertinggi
	fmt.Println("5 buku dengan rating tertinggi:")
	limit := 5
	if n < limit {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("Judul: %s, Rating: %d\n", pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	// Pencarian buku berdasarkan rating dengan binary search
	left, right := 0, n-1
	var found bool
	for left <= right {
		mid := left + (right-left)/2
		if pustaka[mid].rating == r {
			fmt.Printf("Buku ditemukan dengan rating %d: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
				pustaka[mid].rating, pustaka[mid].judul, pustaka[mid].penulis,
				pustaka[mid].penerbit, pustaka[mid].tahun)
			found = true
			break
		} else if pustaka[mid].rating < r {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	DaftarkanBuku(&pustaka, &n)
	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka, n)
	// Cetak buku favorit dengan rating tertinggi
	CetakTerfavorit(pustaka, n)
	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka, n)

	var ratingCari int
	fmt.Println("Masukkan rating yang ingin dicari:")
	fmt.Scanf("%d\n", &ratingCari)
	CariBuku(pustaka, n, ratingCari)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%206/ss/sl4.png)
