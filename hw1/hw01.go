package main

import (
	"fmt"
	"log"
)

func main() {
	var destination string
	var lightYears float64
	var speed float64
	var sLight = 186000.0
	var yearDays = 365.25

	fmt.Printf("Where are you going? ")
	fmt.Scan(&destination)

	fmt.Printf("How far away is %s (light years)? ", destination)
	_, err := fmt.Scan(&lightYears)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("How fast is your spaceship (miles per hour)? ")
	_, err2 := fmt.Scan(&speed)
	if err2 != nil {
		log.Fatal(err2)
	}

	var hourInYear = yearDays * 24.0
	var sLightHr = sLight * 3600
	var miles = sLightHr * hourInYear * lightYears

	var hours = miles / speed
	var minutes = hours * 60
	var seconds = minutes * 60
	var days = hours / 24
	var years = days / 365.25
	var centuries = years / 100
	fmt.Printf("Travel Time to %s, (%.2f lightyears, %.2f mph): \n %.2f seconds \n %.2f minutes \n %.2f hours \n %.2f days \n %.2f years \n %.2f centuries", destination, lightYears, speed, seconds, minutes, hours, days, years, centuries)
}
