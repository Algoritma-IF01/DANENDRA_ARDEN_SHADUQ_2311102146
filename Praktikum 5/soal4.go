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
