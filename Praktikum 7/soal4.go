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
