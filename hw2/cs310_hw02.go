package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var boolPrintedMes bool
	var boolPrime bool
	var boolFactors bool

	for boolPrime == false {
		rand.Seed(time.Now().UnixNano())
		var randVal int = (rand.Intn(10000000-2) + 2)
		boolFactors = false
		boolPrintedMes = false
		for x := 2; x < randVal; x++ {
			if randVal%x == 0 {
				if !boolPrintedMes {
					fmt.Printf("%d is NOT prime.  factors = ", randVal)
					boolPrintedMes = true
				}
				fmt.Printf("%d ", x)
				boolFactors = true
			}
		}
		if !boolFactors {
			boolPrime = true
			fmt.Printf("%d IS PRIME.", randVal)
		}
		fmt.Print("\n\n")
	}

}
