package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func send(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func FtoC(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()["value"]
	if len(params) == 0 {
		message := "ERROR MISSING GET PARAMETER \"value\" \nusage: http://localhost:5000/FtoC?value=10"
		send(writer, message)
	} else {
		farenheit := params[0]
		f, err := strconv.ParseFloat(farenheit, 64)
		if err != nil {
			//log.Fatal(err)
			message := "ERROR: BAD VALUE\nusage: http://localhost:5000/FtoC?value=10"
			send(writer, message)
		} else {
			celcius := fmt.Sprintf("%.3f F => %.3f C", f, (f-32)*5/9)
			send(writer, celcius)
		}
	}
}

func MtoK(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()["value"]
	if len(params) == 0 {
		message := "ERROR MISSING GET PARAMETER \"value\" \nusage: http://localhost:5000/MtoK?value=10"
		send(writer, message)
	} else {
		miles := params[0]
		m, err := strconv.ParseFloat(miles, 64)
		if err != nil {
			//log.Fatal(err)
			message := "ERROR: BAD VALUE\nusage: http://localhost:5000/MtoK?value=10"
			send(writer, message)
		} else {
			var feet = m * 5280
			var inches = feet * 12
			var cm = inches * 2.54

			kilometers := fmt.Sprintf("%.3f M => %.3f K", m, cm/100000)
			send(writer, kilometers)
		}
	}
}

func GtoL(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()["value"]
	if len(params) == 0 {
		message := "ERROR MISSING GET PARAMETER \"value\" \nusage: http://localhost:5000/GtoL?value=10"
		send(writer, message)
	} else {
		gallons := params[0]
		g, err := strconv.ParseFloat(gallons, 64)
		if err != nil {
			//log.Fatal(err)
			message := "ERROR: BAD VALUE\nusage: http://localhost:5000/GtoL?value=10"
			send(writer, message)
		} else {
			var l = g * 3.785411784

			liters := fmt.Sprintf("%.3f G => %.3f L", g, l)
			send(writer, liters)
		}
	}
}

func main() {
	http.HandleFunc("/FtoC", FtoC)
	http.HandleFunc("/MtoK", MtoK)
	http.HandleFunc("/GtoL", GtoL)
	err := http.ListenAndServe("localhost:5000", nil)
	log.Fatal(err)
}
