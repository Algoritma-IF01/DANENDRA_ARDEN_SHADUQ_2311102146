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
