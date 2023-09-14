package main

import (
	"bufio"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal(errors.New("usage: ./hw05 fileName seed"))
	}

	fileName := os.Args[1]

	var seed int64

	seed, intErr := strconv.ParseInt(os.Args[2], 0, 64)
	if intErr != nil {
		log.Fatal("Error parsing seed")
	}

	rand.Seed(seed)

	FileWords := make([]string, 0)

	iFile, fileError := os.Open(fileName)
	if fileError != nil {
		log.Fatal("error opening file")
	}
	scanner := bufio.NewScanner(iFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		FileWords = append(FileWords, word)
	}

	var randIndex int = (rand.Intn(len(FileWords)))
	wordOne := FileWords[randIndex]
	wordOne = strings.Title(wordOne)

	var equal bool
	randIndex = (rand.Intn(len(FileWords)))
	wordTwo := FileWords[randIndex]
	wordTwo = strings.Title(wordTwo)
	if wordTwo == wordOne {
		equal = true
	}
	for equal {
		randIndex = (rand.Intn(len(FileWords)))
		wordTwo := FileWords[randIndex]
		wordTwo = strings.Title(wordTwo)
		if wordTwo != wordOne {
			equal = false
		}
	}

	var password string
	password = password + wordOne + wordTwo

	var numberPortion string
	for i := 0; i < 5; i++ {
		randIndex = (rand.Intn(10))
		numberPortion = numberPortion + strconv.Itoa(randIndex)
	}

	password = password + numberPortion

	symbols := [8]string{"!", "@", "#", "$", "%", "^", "&", "*"}
	randIndex = (rand.Intn(8))

	symbolC := symbols[randIndex]
	password = password + symbolC
	fmt.Printf("plaintext password = %s \n", password)

	h := sha1.New()
	h.Write([]byte(password))
	fmt.Printf("SHA1 password hash = %x", h.Sum(nil))
}
