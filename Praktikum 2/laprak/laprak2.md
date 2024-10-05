# <h1 align="center">PRAKTIKUM 2</h1>
<p align="center">DANENDRA ARDEN SHADUQ - 2311102146</p>

### 1. hello.go
```GO
package main
import "fmt"

func main() {
	var greetings = "Selamat datang di dunia DAP"
	var a, b int
	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
```

Output
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%202/ss/1.png)

### 2. Hipotenusa
```GO
package main
import "fmt"

func main() {
	var a, b, c float64
	var hipotenusa bool
	fmt.Print("Masukkan nilai A : ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai B : ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan nilai C : ")
	fmt.Scanln(&c)

	fmt.Scanln(&a, &b, &c)
	hipotenusa = (c*c) == (a*a + b*b)
	fmt.Println("Sisi c adalah hipotenusa segitiga a,b,c: ", hipotenusa)
}
```
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%202/ss/2.png)

### 3. Latihan 1
```GO
package main
import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Println("Masukkan input string : ")
	fmt.Scanln(&satu)
	fmt.Println("Masukkan input string : ")
	fmt.Scanln(&dua)
	fmt.Println("Masukkan input string : ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%202/ss/3.png)

### 4. Latihan 2
```GO
package main
import "fmt"

func main() {
	var tahun int

	fmt.Print("Masukkan Tahun : ")
	fmt.Scanln(&tahun)

	if tahun%4==0{
		fmt.Println("Tahun kabisat : true")
	}else{
		fmt.Println("Tahun kabisat : false")
	}
}
```
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%202/ss/4.png)

### 5. Latihan 3
```GO
package main

import (
	"fmt"
)

func main() {
	var jariJari float64
	fmt.Print("Jejari = ")
	fmt.Scanln(&jariJari)

	// Menghitung volume
	volume := (4.0 / 3.0) * 3.1415926535 * jariJari * jariJari * jariJari

	// Menghitung luas permukaan
	luas := 4 * 3.1415926535 * jariJari * jariJari

	// Menampilkan hasil
	fmt.Printf("Bola dengan jejari %.4f memiliki volume %.4f dan luas kulit %.4f\n", jariJari, volume, luas)
}
```
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%202/ss/5.png)

### 6. Latihan 4
```GO
package main

import "fmt"

func main() {
    var celsius float64

    fmt.Print("Masukkan suhu dalam derajat Celsius: ")
    fmt.Scanln(&celsius)

    // Konversi ke Fahrenheit
    fahrenheit := (celsius * 9 / 5) + 32

    // Konversi ke Reamur
    reamur := celsius * 4 / 5

    // Konversi ke Kelvin
    kelvin := (fahrenheit + 459.67) * 5 / 9

    fmt.Printf("Suhu dalam derajat Fahrenheit: %.2f\n", fahrenheit)
    fmt.Printf("Suhu dalam derajat Reamur: %.2f\n", reamur)
    fmt.Printf("Suhu dalam derajat Kelvin: %.2f\n", kelvin)
}
```
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%202/ss/6.png)

### 7. Latihan 5
```GO
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

```
![240302_00h00m06s_screenshot](https://github.com/Algoritma-IF01/DANENDRA_ARDEN_SHADUQ_2311102146/blob/main/Praktikum%202/ss/7.png)
