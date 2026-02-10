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
	fmt.Print("Введите исходную валюту (EUR/USD/RUB)")
	fmt.Scan(&startCur)
	fmt.Print("Введите число")
	fmt.Scan(&inputCur)
	fmt.Print("Введите целевую валюту (EUR/USD/RUB)")
	fmt.Scan(&endCur)
	err := checkInput(inputCur, startCur, endCur)
	return inputCur, startCur, endCur, err
}

func checkInput(inputCur float64, startCur string, endCur string) error {
	var err error
	if startCur != "EUR" && startCur != "USD" && startCur != "RUB" {
		err = errors.New("Вы не правильно ввели имя валюты")
	}
	if endCur != "EUR" && endCur != "USD" && endCur != "RUB" {
		err = errors.New("Вы не правильно ввели имя валюты")
	}
	if inputCur < 0 {
		err = errors.New("Неверное число")
	}
	return err
}

func calculateCur(inputCur float64, startCur string, endCur string) float64 {
	var result, course float64
	switch {
	case startCur == "USD" && endCur == "RUB":
		course = 77.05
	case startCur == "USD" && endCur == "EUR":
		course = 0.84
	case startCur == "EUR" && endCur == "USD":
		course = 1.19
	case startCur == "EUR" && endCur == "RUB":
		course = 92.07
	case startCur == "RUB" && endCur == "USD":
		course = 0.013
	case startCur == "RUB" && endCur == "EUR":
		course = 0.011
	}
	result = inputCur * course
	return result
}
