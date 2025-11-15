package main

import "fmt"

const (
	fromUSDtoEUR = 0.86
	fromUSDtoRUB = 81.13
)

func main() {
	fromEURtoRUB := fromUSDtoRUB / fromUSDtoEUR

	fmt.Printf("Конвертация из EUR в RUB: %.2f", fromEURtoRUB)
}
