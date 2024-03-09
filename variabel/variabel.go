package main

import "fmt"

func main() {
	var firstName string = "Fahri"

	var lastName string = "Desrizal"

	fmt.Println(firstName, lastName)
	fmt.Printf("Hallo, Fahri Desrizal\n")
	fmt.Printf("hallo, %s %s\n", firstName, lastName)

	var (
		age     int
		address string
	)

	age = 23
	address = "bogor"

	fmt.Printf("Halo, saya %s %s, saya berumur %d dan sedang berada di %s\n", firstName, lastName, age, address)

	var (
		bootcamptName, bootcampAddress = "enigmacampt", "Jakarta"
	)

	fmt.Println(bootcamptName, bootcampAddress)

	day := "Wednesday"
	date := "13"
	month := "December"

	fmt.Println(day, date, month+"2000")
}
