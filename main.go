package main

import (
	"errors"
	"fmt"
)

// Домашнее задание 1-start

func main() {
	for {
		inputCur, startCur, endCur, err := userInput()
		if err != nil {
			fmt.Print(err)
			continue
		}
		fmt.Println(calculateCur(inputCur, startCur, endCur))
	}
}

func userInput() (float64, string, string, error) {
	var inputCur float64
	var startCur string
	var endCur string
	var err error
	fmt.Print("Введите исходную валюту (EUR/USD/RUB)")
	fmt.Scan(&startCur)
	if startCur != "EUR" || startCur != "USD" || startCur != "RUB" {
		err = errors.New("Вы не правильно ввели имя валюты")
	}
	fmt.Print("Введите число")
	fmt.Scan(&inputCur)
	if inputCur < 0 {
		err = errors.New("Неверное число")
	}
	fmt.Print("Введите целевую валюту (EUR/USD/RUB)")
	fmt.Scan(&endCur)
	if endCur != "EUR" || endCur != "USD" || endCur != "RUB" {
		err = errors.New("Вы не правильно ввели имя валюты")
	}
	return inputCur, startCur, endCur, err
}

func calculateCur(inputCur float64, startCur string, endCur string) float64 {
	const usdToEur float64 = 0.84
	const usdToRub float64 = 77.05
	var eurToRub float64 = inputCur * usdToRub / usdToEur
	return eurToRub
}
