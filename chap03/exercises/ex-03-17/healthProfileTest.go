package main

import (
	"fmt"

	hp "github.com/biraneves/go-exercises/chap03/exercises/ex-03-17/health-profile"
	sv "github.com/biraneves/screen-visual"
)

func main() {

	var firstName string
	fmt.Print("First name: ")
	fmt.Scanln(&firstName)

	var lastName string
	fmt.Print("Last name: ")
	fmt.Scanln(&lastName)

	var sex string
	fmt.Print("Sex (F/M): ")
	fmt.Scanln(&sex)

	var birthDay int
	fmt.Print("Birth day (1 - 31): ")
	fmt.Scanln(&birthDay)

	var birthMonth int
	fmt.Print("Birth month (1 - 12): ")
	fmt.Scanln(&birthMonth)

	var birthYear int
	fmt.Print("Birth year (4 digits): ")
	fmt.Scanln(&birthYear)

	var height float64
	fmt.Print("Height (m): ")
	fmt.Scanln(&height)

	var weight float64
	fmt.Print("Weight (kg): ")
	fmt.Scanln(&weight)

	health := hp.NewHealthProfile(firstName, lastName, sex, birthDay, birthMonth, birthYear, height, weight)
	currentYear := 2022

	sv.Lineln("=", 64)

	fmt.Printf("Name: %s\n", health.GetFullName())
	fmt.Printf("Sex: %s\n", health.Sex)
	fmt.Printf("Birth date: %s\n", health.GetBirthDate())
	fmt.Printf("Age (approx.): %d\n", health.GetAge(currentYear))
	fmt.Printf("Height: %.1fm\n", health.GetHeight())
	fmt.Printf("Weight: %.1fkg\n", health.GetWeight())

	sv.Lineln("-", 64)

	fmt.Printf("Max heart rate: %d bpm\n", health.GetMaxHeartRate(currentYear))
	minHR, maxHR := health.GetTargetHeartRate(currentYear)
	fmt.Printf("Target heart rate: (%d - %d) bpm\n", minHR, maxHR)
	fmt.Printf("BMI: %.1fkg/m^2\n", health.GetBMI())

	sv.Lineln("=", 64)

	fmt.Println("")
	fmt.Print(health.GetBMITable())

}