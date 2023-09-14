package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func fib(fibNum float64, m map[float64]float64) float64 {
	var value float64
	var ok bool
	if fibNum == 0 {
		//m[fibNum] = 0
		return 0
	} else if fibNum == 1 {
		m[fibNum] = 1
		return 1
	} else {
		value, ok = m[fibNum]
		if ok {
			return value
		} else {
			m[fibNum] = fib(fibNum-1, m) + fib(fibNum-2, m)
			return fib(fibNum-1, m) + fib(fibNum-2, m)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(errors.New("usage: ./hw06 N"))
	}

	N, intErr := strconv.ParseFloat(os.Args[1], 64)
	if intErr != nil {
		log.Fatal("Error with setting nth fib number")
	}
	if N < 0 || N > 1000 {
		log.Fatal("range of N: 0 <= N <= 1000")
	}
	//fmt.Print(N)
	my_map := make(map[float64]float64)

	result := fib(N, my_map)

	//fmt.Println(my_map)
	fibNumSlice := make([]float64, 0)
	for _, v := range my_map {
		fibNumSlice = append(fibNumSlice, v)
	}
	sort.Float64s(fibNumSlice)

	fmt.Printf("fibb(%.0f) = %.0f \n", N, result)

	for k, v := range fibNumSlice {
		var x int = int(N)
		if k < x-1 {
			fmt.Printf("fibb(%d) =  %.0f \n", k+1, v)
		}
	}
}
