package main

import "fmt"

func main() {
	var n float64
	var e float64

	fmt.Scan(&n)
	var auxE float64 = 1
	var factorial float64 = 1
	var aux float64 = 1
	var i float64
	for i = 1; i <= n; i++ {
		factorial = aux * i
		aux = factorial
		e = (1.0 / factorial) + auxE
		auxE = e

	}
	fmt.Println(e)
}
