package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Node struct {
	name          string
	adjacentNodes map[string]Node
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal(errors.New("usage: ./hw07 fileName"))
	}

	fileName := os.Args[1]

	overallMap := make(map[string]Node)
	combos := make(map[string]int)
	var n1 Node
	var n2 Node

	iFile, fileError := os.Open(fileName)
	if fileError != nil {
		log.Fatal("error opening file")
	}
	scanner := bufio.NewScanner(iFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		splitString := strings.Split(line, " -> ")
		if len(splitString) != 2 {
			log.Fatal("error in file format: COLOR_# -> COLOR_#")
		}
		_, ok := overallMap[splitString[0]]
		if !ok {
			n1 = Node{name: splitString[0], adjacentNodes: make(map[string]Node)}
			overallMap[splitString[0]] = n1
		} else {
			n1 = overallMap[splitString[0]]
		}

		_, check2 := overallMap[splitString[1]]
		if !check2 {
			n2 = Node{name: splitString[1], adjacentNodes: make(map[string]Node)}
			overallMap[splitString[1]] = n2
			n1.adjacentNodes[splitString[1]] = n2
		} else if check2 {
			n2 = overallMap[splitString[1]]
			n1.adjacentNodes[splitString[1]] = n2
		}

		color_one := strings.Split(splitString[0], "_")
		color1 := color_one[0]
		color_two := strings.Split(splitString[1], "_")
		color2 := color_two[0]
		colorCode := color1 + " -> " + color2

		value, colorCheck := combos[colorCode]
		if colorCheck {
			value++
			combos[colorCode] = value
		} else {
			combos[colorCode] = 1
		}
	}

	keySlice := make([]string, 0)

	for k, _ := range overallMap {
		//fmt.Println(k)
		keySlice = append(keySlice, k)
	}

	sort.Strings(keySlice)

	fmt.Printf("==== GRAPH ====")
	for _, v := range keySlice {
		var index Node = overallMap[v]
		connectionSlice := make([]string, 0)
		for j, _ := range index.adjacentNodes {
			connectionSlice = append(connectionSlice, j)
		}
		sort.Strings(connectionSlice)
		fmt.Printf("\n %s", v)
		for _, y := range connectionSlice {
			fmt.Printf("\n\t==> %s", y)
		}
	}

	fmt.Printf("\n==== COMBOS ==== \n")

	colorSlice := make([]string, 0)
	for c, _ := range combos {
		colorSlice = append(colorSlice, c)
	}

	sort.Strings(colorSlice)
	for _, res := range colorSlice {
		fmt.Printf("%s = %d\n", res, combos[res])
	}
}
