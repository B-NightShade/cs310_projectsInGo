package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type DataEnc struct {
	Text  string
	Shift string
}

type DataDec struct {
	Cipher string
	Shift  string
}

func encrypt(writer http.ResponseWriter, request *http.Request) {
	var params []string
	text := request.URL.Query()["text"]
	shift := request.URL.Query()["shift"]
	strText := strings.Join(text, "")
	shiftText := strings.Join(shift, "")
	params = append(params, strText)
	params = append(params, shiftText)
	data := DataEnc{Text: strText, Shift: shiftText}

	if request.Method == "GET" {
		if len(params) == 0 {
			html, errorEnc := template.ParseFiles("encrypt.html")
			if errorEnc != nil {
				fmt.Printf("ERROR: %x", errorEnc)
			}
			html.Execute(writer, nil)
		} else {
			html, errorEnc := template.ParseFiles("encrypt.html")
			if errorEnc != nil {
				fmt.Printf("ERROR: %x", errorEnc)
			}
			html.Execute(writer, data)
		}
	}
	if request.Method == "POST" {
		textForm := request.FormValue("Text")
		//fmt.Println(textForm)
		shiftForm := request.FormValue("shift")
		shift, errAtoI := strconv.Atoi(shiftForm)
		if errAtoI != nil {
			fmt.Println(errAtoI)
		}
		s := rune(shift)
		//fmt.Print(shiftForm)

		chars := []rune(textForm)
		for i, c := range chars {
			c1 := ((c-'a')+s)%26 + 'a'
			chars[i] = c1
		}
		cipher := string(chars)
		//fmt.Println(cipher)
		http.Redirect(writer, request, "/decrypt?cipher="+cipher+"&shift="+shiftForm, http.StatusFound)
	}
}

func decrypt(writer http.ResponseWriter, request *http.Request) {
	var params []string
	cipher := request.URL.Query()["cipher"]
	shift := request.URL.Query()["shift"]
	cipherText := strings.Join(cipher, "")
	shiftText := strings.Join(shift, "")
	params = append(params, cipherText)
	params = append(params, shiftText)
	data := DataDec{Cipher: cipherText, Shift: shiftText}

	if request.Method == "GET" {
		if len(params) == 0 {
			html, errorEnc := template.ParseFiles("decrypt.html")
			if errorEnc != nil {
				fmt.Printf("ERROR: %x", errorEnc)
			}
			html.Execute(writer, nil)
		} else {
			html, errorEnc := template.ParseFiles("decrypt.html")
			if errorEnc != nil {
				fmt.Printf("ERROR: %x", errorEnc)
			}
			html.Execute(writer, data)
		}
	}
	if request.Method == "POST" {
		CipherForm := request.FormValue("Text")
		shiftForm := request.FormValue("shift")
		shift, errAtoI := strconv.Atoi(shiftForm)
		if errAtoI != nil {
			fmt.Println(errAtoI)
		}
		s := rune(shift)

		chars := []rune(CipherForm)
		for i, c := range chars {
			c1 := ((c-'z')-s)%26 + 'z'
			chars[i] = c1
		}
		text := string(chars)
		http.Redirect(writer, request, "/encrypt?text="+text+"&shift="+shiftForm, http.StatusFound)
	}

}

func main() {
	http.HandleFunc("/encrypt", encrypt)
	http.HandleFunc("/decrypt", decrypt)
	err := http.ListenAndServe("localhost:5000", nil)
	log.Fatal(err)
}
