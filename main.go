package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	fromUSDtoEUR = 0.86
	fromUSDtoRUB = 81.13
	fromEURtoRUB = fromUSDtoRUB / fromUSDtoEUR
)

func getDataFromUser() (float64, string, string) {
	var (
		amount       float64
		currencyFrom string
		currencyTo   string
	)
	var currencySet []string = []string{"rub", "usd", "eur"}
	availableCurrencySetTo := make([]string, 0, len(currencySet)-1)

	currencySetString := convertArrToString(currencySet)

	// Ввод первой валюты
	fmt.Print("Введите валюту для конвертации: ", currencySetString)

	for {
		_, err := fmt.Scan(&currencyFrom)
		currencyFrom = strings.ToLower(currencyFrom)
		if err != nil {
			log.Fatalln(err)
		}

		if inputCompare(currencyFrom, currencySet) {
			fmt.Printf("Вы ввели : %s \n", currencyFrom)
			break
		} else {
			fmt.Printf(" Данная валюта не поддерживается калькулятором введите валюту из списка : %s \n", currencySetString)
		}
	}

	fmt.Println("Введите количество единиц :  ")

	_, err := fmt.Scan(&amount)
	if err != nil {
		log.Fatalln(err)

	}

	// Ввод второй валюты
	//  избавимся от валюты которую мы выбрали из списка и получим новый срез без этого значения
	for _, value := range currencySet {
		if value != currencyFrom {
			availableCurrencySetTo = append(availableCurrencySetTo, value)
		}

	}
	availableCurrencySetString := convertArrToString(availableCurrencySetTo)
	fmt.Printf("Введите валюту назначения  : %s\n", availableCurrencySetString)
	for {
		_, err = fmt.Scan(&currencyTo)
		currencyTo = strings.ToLower(currencyTo)
		if err != nil {
			log.Fatalln(err)
		}
		if !inputCompare(currencyTo, availableCurrencySetTo) {
			fmt.Printf("Ввод не верен. Введите : %s\n", availableCurrencySetString)
			continue
		} else {
			break
		}
	}

	return amount, currencyFrom, currencyTo
}
func inputCompare(currencyFrom string, currencySet []string) bool {
	for _, currency := range currencySet {
		if currency == currencyFrom {
			return true
		}
	}
	return false
}

func convertArrToString(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	result := arr[0]

	for i := 1; i < len(arr); i++ {
		result += ", " + arr[i]
	}
	return result
}

func counter(amount float64, currencyFrom string, currencyTo string) float64 {
	fmt.Printf("Данные counter : %v %s %s \n", amount, currencyFrom, currencyTo)
	switch {
	case currencyFrom == "rub" && currencyTo == "usd":
		return amount / fromUSDtoRUB
	case currencyFrom == "rub" && currencyTo == "eur":
		return amount / fromEURtoRUB
	case currencyFrom == "usd" && currencyTo == "eur":
		return amount * fromUSDtoEUR
	case currencyFrom == "usd" && currencyTo == "rub":
		return amount * fromUSDtoRUB
	case currencyFrom == "eur" && currencyTo == "usd":
		return amount / fromUSDtoEUR
	case currencyFrom == "eur" && currencyTo == "rub":
		return amount * fromEURtoRUB
	default:
		return 0.0
	}
}

func main() {
	amount, currencyFrom, currencyTo := getDataFromUser()
	response := counter(amount, currencyFrom, currencyTo)
	fmt.Printf("Конвертация %v единиц из %s в %s результат : %.2f\n", amount, currencyFrom, currencyTo, response)
}
