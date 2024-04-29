package main

import (
	"fmt"
)

func main() {
	var firstName string = "Alfaathir"
	var lastName string = "Rasyid"

	fmt.Println(firstName)
	fmt.Printf("Hallo, nama saya %s %s\n", firstName, lastName)

	var (
		age     int
		address string
	)
	age = 17
	address = "Kendari"

	fmt.Printf("Hallo nama saya %s %s, saya berumur %d tahun dan saya tinggal di %s\n", firstName, lastName, age, address)

	var (
		bootcampName, bootcampAddress = "Enigma Camp", "Jakrta selatan"
	)
	fmt.Println(bootcampName, bootcampAddress)

	day := "Monday"
	date := 31
	mounth := "October"
	fmt.Printf("%s %d %s 2024\n", day, date, mounth)

	const phi = 3.14

	bootcamp, _ := "enigma camp", "aktif7"

	fmt.Println(bootcamp)
	fmt.Println(phi)
}
