package main

import (
	"fmt"
	"log"
)

const (
	fromUSDtoEUR = 0.86
	fromUSDtoRUB = 81.13
)

func getDataFromUser() (int, int, int) {
	var (
		amount       int
		currencyFrom int
		currencyTo   int
	)
	_, err := fmt.Scan(&amount)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Scan(&currencyFrom)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Scan(&currencyTo)
	if err != nil {
		log.Fatalln(err)
	}

	return amount, currencyFrom, currencyTo
}
func counter(amount int, currencyFrom int, currencyTo int) int {
	return 0
}

func main() {
	fromEURtoRUB := fromUSDtoRUB / fromUSDtoEUR

	fmt.Printf("Конвертация из EUR в RUB: %.2f", fromEURtoRUB)
}
