# <h1 align="center">PRAKTIKUM 3</h1>
<p align="center">DANENDRA ARDEN SHADUQ - 2311102146</p>

## PERULANGAN
### 1. Latihan 1
```GO
package main

import (
	"fmt"
)

func main() {
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d\n", i)
		fmt.Print("Masukkan warna pertama : ")
		fmt.Scan(&warna1)
		fmt.Print("Masukkan warna kedua : ")
		fmt.Scan(&warna2)
		fmt.Print("Masukkan warna ketiga : ")
		fmt.Scan(&warna3)
		fmt.Print("Masukkan warna keempat : ")
		fmt.Scan(&warna4)

		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] || warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {
		hasil = false
		}
	}

	fmt.Println("BERHASIL : ", 	hasil)
}
```

Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%203/ss/L1.png)

### 2. Latihan 2
```GO
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var pita string
	var bungaCount int

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita += " - " + input
		}
		bungaCount++
	}
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", bungaCount)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%203/ss/L2.png)

### 3. Tugas 1
```GO
package main

import (
	"fmt"
)

func main() {
	var kantong1, kantong2 float64

	for {
		// Input
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		// Stop the process if one of the bag weights is negative
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Calculate the weight difference
		diff := kantong1 - kantong2
		if diff < 0 {
			diff = -diff // convert to positive
		}

		// Determine if the motorcycle will lose balance
		if diff >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}

		// Stop the process if the total weight of both bags together exceeds or equals 150kg
		totalWeight := kantong1 + kantong2
		if totalWeight >= 150 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}

```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%203/ss/T1.png)

### 4. Tugas 2
```GO
package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	// Initialize the product for sqrt(2) approximation
	approximation := 1.0

	// Loop through k from 0 to K and compute the product
	for k := 0; k <= K; k++ {
		numerator := (4*float64(k) + 2) * (4*float64(k) + 2) // (4k + 2)^2
		denominator := (4*float64(k) + 1) * (4*float64(k) + 3) // (4k + 1)(4k + 3)
		approximation *= numerator / denominator
	}

	// Output the result, approximating sqrt(2)
	fmt.Printf("Nilai akar 2 = %.10f\n", approximation)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%203/ss/T2.png)

## PERCABANGAN
### 1. Tugas 1
```GO
package main

import (
	"fmt"
)

func main() {
	const biayaPerKg = 10000
	const biayaTambahanDiAtas500Gram = 5
	const biayaTambahanDiBawah500Gram = 15

	var beratDalamGram int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratDalamGram)

	// Konversi berat ke kg dan gram
	kg := beratDalamGram / 1000
	gram := beratDalamGram % 1000

	// Hitung biaya dasar
	biayaDasar := kg * biayaPerKg

	// Hitung biaya tambahan
	var biayaTambahan int
	if beratDalamGram > 10000 {
		biayaTambahan = 0
	} else if gram >= 500 {
		biayaTambahan = gram * biayaTambahanDiAtas500Gram
	} else {
		biayaTambahan = gram * biayaTambahanDiBawah500Gram
	}

	// Hitung total biaya
	totalBiaya := biayaDasar + biayaTambahan

	// Buat detail output
	detailBerat := fmt.Sprintf("%d kg + %d gr", kg, gram)
	detailBiaya := fmt.Sprintf("Rp. %d + Rp. %d", biayaDasar, biayaTambahan)

	fmt.Printf("Berat parsel (gram): %d\nDetail berat: %s\nDetail biaya: %s\nTotal biaya: Rp. %d\n", beratDalamGram, detailBerat, detailBiaya, totalBiaya)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%203/ss/T-1.png)

### 2. Tugas 2A
```GO
package main
import "fmt"
func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam >= 80 {
		nam = "A"
	}
	if nam >= 72.5 {
		nam = "AB"
	}
	if nam > 65 {
		nam = "B"
	}
	if nam > 57.5 {
		nam = "BC"
	}
	if nam > 50 {
		nam = "C"
	}
	if nam > 40 {
		nam = "D"
	} else if nam <= 40 {
		nam = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%203/ss/T-2A.png)

### 3. Tugas 2B
Apa saja kesalahan dari program tersebut? Mengapa demikian?

1. Tipe Data yang Tidak Sesuai
   Deklarasi variabel `nam` sebagai `float64`, tetapi digunakan pada nilai string (misalnya "A", "B") ke dalamnya.
   
2. Urutan Kondisi yang Salah
   Kondisi `if` tidak ditulis dalam urutan yang benar. Misalnya, jika `nam` lebih besar atau sama dengan 80, maka kondisi untuk nilai lainnya (seperti 72.5, 65, 57.5) juga akan terpenuhi, menyebabkan hasil yang tidak diinginkan.

3. Variabel yang Tidak Digunakan
   Variabel `nmk` dideklarasikan tetapi tidak digunakan untuk menyimpan nilai huruf mata kuliah.

Jelaskan alur program seharusnya!

1. Program dimulai dan variabel `nam` dan `nmk` dideklarasikan.
2. Program meminta pengguna untuk memasukkan nilai akhir mata kuliah dan membaca nilai tersebut menggunakan `fmt.Scanln`.
3. Program mengevaluasi nilai akhir mata kuliah (`nam`) dan menetapkan nilai huruf ke variabel `nmk` berdasarkan rentang nilai:
   - Jika `nam >= 80`, maka `nmk` diatur ke "A".
   - Jika `nam >= 72.5`, maka `nmk` diatur ke "AB".
   - Jika `nam > 65`, maka `nmk` diatur ke "B".
   - Jika `nam > 57.5`, maka `nmk` diatur ke "BC".
   - Jika `nam > 50`, maka `nmk` diatur ke "C".
   - Jika `nam > 40`, maka `nmk` diatur ke "D".
   - Jika `nam <= 40`, maka `nmk` diatur ke "E".
4. Program mencetak nilai huruf mata kuliah (`nmk`).

### 4. Tugas 2C Perbaikan Program
```GO
package main

import "fmt"

func main() {
    var nam float64
    var nmk string
    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scanln(&nam)

    if nam >= 80 {
        nmk = "A"
    } else if nam >= 72.5 {
        nmk = "AB"
    } else if nam > 65 {
        nmk = "B"
    } else if nam > 57.5 {
        nmk = "BC"
    } else if nam > 50 {
        nmk = "C"
    } else if nam > 40 {
        nmk = "D"
    } else {
        nmk = "E"
    }

    fmt.Println("Nilai mata kuliah: ", nmk)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%203/ss/T-2C.png)

### 4. Tugas 3
```GO
package main

import "fmt"

func main() {
    var angka int
    fmt.Print("Bilangan: ")
    fmt.Scan(&angka)

    // Mencari faktor
    factors := []int{}
    for i := 1; i <= angka; i++ {
        if angka%i == 0 {
            factors = append(factors, i)
        }
    }
    fmt.Println("Faktor:", factors)

    // Mengecek prima
    isPrime := true
    for i := 2; i*i <= angka; i++ {
        if angka%i == 0 {
            isPrime = false
            break
        }
    }
    fmt.Printf("Prima: %t\n", isPrime)
}
```
Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%203/ss/T-3.png)
