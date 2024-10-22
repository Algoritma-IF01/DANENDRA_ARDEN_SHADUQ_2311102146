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