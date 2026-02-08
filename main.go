package main

import "fmt"

func main() {

	var eur float64 = 1
	const usdToEur float64 = 0.84
	const usdToRub float64 = 77.05
	var eurToRub float64 = eur * usdToRub / usdToEur
	fmt.Println(eurToRub)

}
