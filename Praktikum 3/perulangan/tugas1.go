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
