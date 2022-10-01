package main

import (
	"fmt"
	"time"
	"github.com/MatWich/Console-clock/digits"
)

func main() {

	numbers := digits.GetDigits()
	for line := range numbers[0] {
		for digit := range numbers {
			fmt.Print(numbers[digit][line], " ")
		}
		fmt.Println()
	}

	currentTime := time.Now().Local()

	fmt.Println(currentTime)
	hour, min, sec := currentTime.Hour(), currentTime.Minute(), currentTime.Second()
	fmt.Printf("%d:%d:%d\n", hour, min, sec)

	colon := digits.GetColon()

	clock := [...][5]string {
		numbers[hour/10], 
		numbers[hour%10], 
		colon,
		numbers[min/10], 
		numbers[min%10],
		colon,
		numbers[sec/10],
		numbers[sec%10],

	}

	for line := range clock[0] {
		for digit := range clock {
			fmt.Print(clock[digit][line])
		}
		fmt.Println()
	}

}
