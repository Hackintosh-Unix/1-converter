package main

import "fmt"

// Домашнее задание 1-start

func main() {

	const usdToEur float64 = 0.84
	const usdToRub float64 = 77.05
	const eurToRub float64 = 1.0 * usdToRub / usdToEur
	fmt.Println(eurToRub)

}
