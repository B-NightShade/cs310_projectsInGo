package main

import (
	"fmt"
	"log"
)

func getOption() string {
	var option string
	fmt.Println("=== Options ====")
	fmt.Println("F -> Fahrenheit to Celsius")
	fmt.Println("M -> Miles to Kilometers")
	fmt.Println("P -> Pounds to Kilograms")
	fmt.Println("Q -> QUIT")
	fmt.Println("===========")
	fmt.Print("ENTER YOUR CHOICE: ")
	fmt.Scan(&option)
	return option
}

func getNum() (float64, error) {
	var enteredNum float64
	fmt.Print("ENTER DECIMAL TO CONNVERT: ")
	_, err := fmt.Scan(&enteredNum)
	if err != nil {
		return 0, err
	} else {
		return enteredNum, nil
	}
}

func printResultHeader() {
	fmt.Println(">>>>>> RESULTS <<<<<<")
}
func printResultFooter() {
	fmt.Println(">>>>>>>>>><<<<<<<<<<")
}

func ftoC(original float64) {
	var Celcius = (original - 32) * 5 / 9
	printResultHeader()
	fmt.Printf("%f degrees F => %f degrees C \n", original, Celcius)
	printResultFooter()
}

func miToKm(original float64) {
	var feet = original * 5280
	var inches = feet * 12
	var cm = inches * 2.54
	var Km = cm / 100000
	printResultHeader()
	fmt.Printf("%f miles => %f km \n", original, Km)
	printResultFooter()
}

func pToK(original float64) {
	var kilograms = original * 0.45359237
	printResultHeader()
	fmt.Printf("%f pounds => %f kg \n", original, kilograms)
	printResultFooter()
}

func main() {
	for true {
		var cOption string
		var numConvert float64
		var validNum = false
		var validOption = false

		cOption = getOption()
		for !validOption {
			if cOption[0:1] != "F" && cOption[0:1] != "M" && cOption[0:1] != "P" && cOption[0:1] != "Q" {
				fmt.Printf("INVALID OPTION: %s \n", cOption)
				cOption = getOption()
			} else {
				validOption = true
			}
		}

		if cOption[0:1] == "Q" {
			break
		}

		numConvert, okNum := getNum()
		for !validNum {
			if okNum != nil {
				log.Print(okNum)
				numConvert, okNum = getNum()
			} else {
				validNum = true
			}
		}

		if cOption[0:1] == "F" {
			ftoC(numConvert)
		} else if cOption[0:1] == "M" {
			miToKm(numConvert)
		} else {
			pToK(numConvert)
		}
	}
}
