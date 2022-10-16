package main

import (
	"fmt"
	"github.com/MatWich/Console-clock/digits"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	numbers := digits.GetDigits()
	for line := range numbers[0] {
		for digit := range numbers {
			fmt.Print(numbers[digit][line], " ")
		}
		fmt.Println()
	}
	// fmt.Println("\033[2J")

	screen.Clear()

	for {
		screen.MoveTopLeft()
		currentTime := time.Now().Local()

		fmt.Println(currentTime)
		hour, min, sec := currentTime.Hour(), currentTime.Minute(), currentTime.Second()
		fmt.Printf("%d:%d:%d\n", hour, min, sec)

		colon := digits.GetColon()

		clock := [...]digits.Placeholder{
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
				fmt.Print(clock[digit][line], " ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}

}
