package main

import (
	"fmt"
	"github.com/MatWich/Console-clock/myvars"
)

func main() {

	digits := myvars.GetDigits()
	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line], " ")
		}
		fmt.Println()
	}
}
