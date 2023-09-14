package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fileName string
	fmt.Print("enter file name: ")
	fmt.Scan(&fileName)
	fmt.Print("\n")

	iFile, fileOpenErr := os.Open(fileName)
	if fileOpenErr != nil {
		log.Fatal(fileOpenErr)
	} else {
		scanner := bufio.NewScanner(iFile)
		scanner.Split(bufio.ScanLines)
		scanner.Scan() //skip to next token
		for scanner.Scan() {
			line := scanner.Text()
			lineParse := strings.Split(line, ",")

			name := lineParse[0]

			yards, yrdsErr := strconv.ParseFloat(lineParse[1], 64)
			if yrdsErr != nil {
				log.Fatal(yrdsErr)
			}

			attempts, attErr := strconv.ParseFloat(lineParse[2], 64)
			if attErr != nil {
				log.Fatal(attErr)
			}

			completions, compErr := strconv.ParseFloat(lineParse[3], 64)
			if compErr != nil {
				log.Fatal(compErr)
			}

			touchdowns, touchErr := strconv.ParseFloat(lineParse[4], 64)
			if touchErr != nil {
				log.Fatal(touchErr)
			}

			interceptions, interErr := strconv.ParseFloat(lineParse[5], 64)
			if interErr != nil {
				log.Fatal(interErr)
			}

			a := (((completions / attempts) - 0.3) * 5)
			if a < 0 {
				a = 0
			} else if a > 2.375 {
				a = 2.375
			}

			b := (((yards / attempts) - 3) * 0.25)
			if b < 0 {
				b = 0
			} else if b > 2.375 {
				b = 2.375
			}

			c := ((touchdowns / attempts) * 20)
			if c < 0 {
				c = 0
			} else if c > 2.375 {
				c = 2.375
			}

			d := (2.375 - ((interceptions / attempts) * 25))
			if d < 0 {
				d = 0
			} else if d > 2.375 {
				d = 2.375
			}

			rating := (((a + b + c + d) / 6) * 100)
			fmt.Printf("%s : %.1f \n", name, rating)
		}
	}
	iFile.Close()
}
