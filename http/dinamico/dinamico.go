package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	/* 	Year: "2006" "06"
	Month: "Jan" "January" "01" "1"
	Day of the week: "Mon" "Monday"
	Day of the month: "2" "_2" "02"
	Day of the year: "__2" "002"
	Hour: "15" "3" "03" (PM or AM)
	Minute: "4" "04"
	Second: "5" "05"
	AM/PM mark: "PM" */
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
