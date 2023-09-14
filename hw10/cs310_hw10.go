package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
)

func worker(id int, routines int, hash string, words []string, ch chan bool) {
	work := len(words) / routines
	if routines > len(words) {
		work = 1
	}
	start := id * work
	end := (id + 1) * work
	extra := len(words) % routines
	for i := start; i < end && i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if j != i && words[i] != words[j] {
				wordOne := words[i]
				wordTwo := words[j]
				//fmt.Printf("%d: %s\n", id, wordOne+wordTwo)
				h := sha1.New()
				h.Write([]byte(wordOne + wordTwo))
				//fmt.Printf("SHA1 password hash = %x\n", h.Sum(nil))
				shaPass := hex.EncodeToString(h.Sum(nil))
				//fmt.Printf("%s\n", shaPass)
				if shaPass == hash {
					fmt.Printf("found: %s %s", wordOne+wordTwo, shaPass)
					ch <- true
				}
			}
		}
	}
	if extra != 0 && id == routines-1 && routines < len(words) {
		for i := end; i < end+extra; i++ {
			for j := 0; j < len(words); j++ {
				if j != i && words[i] != words[j] {
					wordOne := words[i]
					wordTwo := words[j]
					//fmt.Printf("%d: %s\n", id, wordOne+wordTwo)
					h := sha1.New()
					h.Write([]byte(wordOne + wordTwo))
					//fmt.Printf("SHA1 password hash = %x\n", h.Sum(nil))
					shaPass := hex.EncodeToString(h.Sum(nil))
					//fmt.Printf("%s\n", shaPass)
					if shaPass == hash {
						fmt.Printf("found: %s %s", wordOne+wordTwo, shaPass)
						ch <- true
					}
				}
			}
		}
	}
	ch <- false
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Usage: ./hw10 fileName N SHA1")
	}

	fileName := os.Args[1]
	FileWords := make([]string, 0)

	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if n < 1 || n > 100 {
		log.Fatal("N out of range: 1<=N<=100")
	}

	sha1 := os.Args[3]
	//fmt.Println(sha1)

	iFile, fileOpenErr := os.Open(fileName)
	if fileOpenErr != nil {
		log.Fatal(fileOpenErr)
	} else {
		scanner := bufio.NewScanner(iFile)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			word := scanner.Text()
			FileWords = append(FileWords, word)
		}
	}

	ch := make(chan bool)

	for i := 0; i < n; i++ {
		go worker(i, n, sha1, FileWords, ch)
	}

	found := false

	for i := 0; i < n; i++ {
		x := <-ch
		if x == true {
			found = true
		}
	}

	if !found {
		fmt.Println("Password not found!")
		os.Exit(0)
	}

}
