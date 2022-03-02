package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	printvar := os.Getenv("HWC_PRINTVAR")
	if printvar == "true" {
		for _, v := range os.Environ() {
			fmt.Printf("%s \n", v)
		}
	}

	sleepCounter, _ := strconv.Atoi(os.Getenv("HWC_SLEEPTIMER"))

	// increments sleeptimer, starting from HWC_SLEEPTIMER
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello for the %d time \n", sleepCounter)
		fmt.Printf("Hello for the %d time \n", sleepCounter)
		sleepCounter++
	})

	// Prints environment variables
	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		for _, v := range os.Environ() {
			fmt.Fprintf(w, "%s \n", v)
			fmt.Printf("%s \n", v)
		}

	})
	// Kills the service
	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(5)
	})
	// Kills the service after 60 sec
	http.HandleFunc("/timedkill", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This service will be killed in 60 seconds \n")
		fmt.Printf("This service will be killed in 60 seconds \n")
		time.Sleep(60 * time.Second)
		os.Exit(5)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
