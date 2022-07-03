package main

import (
	"fmt"

	"github.com/biraneves/go-exercises/chap03/exercises/ex-03-16/heart-rate"
	sv "github.com/biraneves/screen-visual"
)

func main() {

	var firstName, lastName string
	var birthDay, birthMonth, birthYear int

	fmt.Print("First name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Last name: ")
	fmt.Scanln(&lastName)
	fmt.Print("Birth day (1 - 31): ")
	fmt.Scanln(&birthDay)
	fmt.Print("Birth month (1 - 12): ")
	fmt.Scanln(&birthMonth)
	fmt.Print("Birth year (4 digits): ")
	fmt.Scanln(&birthYear)

	hr := heartRate.NewHeartRate(firstName, lastName, birthDay, birthMonth, birthYear)
	currentYear := 2022

	sv.Lineln("-", 64)

	fmt.Println("Name:", hr.GetFullName())
	fmt.Println("Birth date:", hr.GetBirthDate())
	fmt.Printf("Age (approx.): %d years old.\n", hr.GetAge(currentYear))

	sv.Lineln("-", 64)
	
	fmt.Printf("Max heart rate: %d bpm\n", hr.GetMaxHeartRate(currentYear))
	min, max := hr.GetTargetHeartRates(currentYear)
	fmt.Printf("Target heart rate (min - max): (%d - %d) bpm\n", min, max)
	
	sv.Lineln("-", 64)

}