package main

import "fmt"

const (
	fromUSDtoEUR = 0.86
	fromUSDtoRUB = 81.13
)

func main() {
	fromEURtoRUB := fromUSDtoRUB / fromUSDtoEUR

	fmt.Printf("Из EUR в RUB: %.2f", fromEURtoRUB)
}
