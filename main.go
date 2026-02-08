package main

import "fmt"

// Домашнее задание 1-start

func main() {

	fmt.Println(calculateCur(userInput()))

}

func userInput() (float64, string, string) {
	var inputCur float64
	var startCur string
	var endCur string
	fmt.Scan(&inputCur, &startCur, &endCur)
	return inputCur, startCur, endCur
}

func calculateCur(inputCur float64, startCur string, endCur string) float64 {
	const usdToEur float64 = 0.84
	const usdToRub float64 = 77.05
	var eurToRub float64 = inputCur * usdToRub / usdToEur
	return eurToRub
}
