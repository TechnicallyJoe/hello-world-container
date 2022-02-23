package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	printvar := os.Getenv("HWC_PRINTVAR")
	if printvar == "true" {
		for _, v := range os.Environ() {
			fmt.Printf("%s \n", v)
		}
	}

	sleepCounter, _ := strconv.Atoi(os.Getenv("HWC_SLEEPTIMER"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello for the %d time \n", sleepCounter)
		fmt.Printf("Hello for the %d time \n", sleepCounter)
		sleepCounter++
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		for _, v := range os.Environ() {
			fmt.Fprintf(w, "%s \n", v)
			fmt.Printf("%s \n", v)
		}

	})

	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(5)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
