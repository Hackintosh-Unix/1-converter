package main

import "fmt"

// Домашнее задание 1-start

func main() {

	inputCur := userInput()

	fmt.Println(calculateCur(inputCur))

}

func userInput() float64 {
	var inputCur float64
	fmt.Scan(&inputCur)
	return inputCur
}

func calculateCur(inputCur float64, startCur string, endCur string) float64 {
	const usdToEur float64 = 0.84
	const usdToRub float64 = 77.05
	var eurToRub float64 = inputCur * usdToRub / usdToEur
	return eurToRub
}
