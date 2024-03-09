package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Numerik")
	var positiveNumber uint8 = 90
	var negativeNumber int8 = -90
	fmt.Printf("Bilangan Positif = %d\n", positiveNumber)
	fmt.Printf("Bilagan Negatif = %d\n", negativeNumber)

	var decimalNumber = 3.90
	fmt.Printf("Bilangan Pecahan = %f\n", decimalNumber)
	fmt.Printf("Bilagan Pecahan = %.2f\n", decimalNumber)

	fmt.Printf("========================================\n")

	fmt.Printf("Boolean\n")
	var exist = true
	fmt.Printf("exist ? %t\n", exist)
}
