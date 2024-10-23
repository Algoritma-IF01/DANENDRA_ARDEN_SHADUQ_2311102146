package main

import (
	"fmt"
)

func isValidVoucher(voucher string) bool {
	length := len(voucher)

	if length != 5 && length != 6 {
		return false
	}

	firstTwo := (int(voucher[0]-'0') * int(voucher[1]-'0'))
	lastTwo := (int(voucher[length-2]-'0') * int(voucher[length-1]-'0'))

	if firstTwo != lastTwo {
		return false
	}

	var middleDigit int
	if length == 5 {
		middleDigit = int(voucher[2] - '0') 
	} else {
		middleDigit = int(voucher[2]-'0') + int(voucher[3]-'0') + int(voucher[4]-'0') 
	}

	if middleDigit%2 != 0 {
		return false
	}

	return true
}

func main() {
	var N int
	fmt.Scan(&N)

	validCount := 0
	invalidCount := 0

	for i := 0; i < N; i++ {
		var voucher string
		fmt.Scan(&voucher)

		if isValidVoucher(voucher) {
			validCount++
		} else {
			invalidCount++
		}
	}

	fmt.Printf("Banyaknya voucher valid: %d\n", validCount)
	fmt.Printf("Banyaknya voucher tidak valid: %d\n", invalidCount)
}

