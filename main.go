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
	timedKill := os.Getenv("HWC_TIMEDKILL")
	sleepTimer, _ := strconv.Atoi(os.Getenv("HWC_SLEEPTIMER"))
	exitcode, _ := strconv.Atoi(os.Getenv("HWC_EXITCODE"))
	if timedKill == "true" {
		fmt.Printf("This service will be killed in %d seconds \n", sleepTimer)
		time.Sleep(time.Duration(sleepTimer) * time.Second)
		os.Exit(exitcode)
	}

	counter := 0
	// increments sleeptimer, starting from HWC_SLEEPTIMEr
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello for the %d time \n", counter)
		fmt.Printf("Hello for the %d time \n", counter)
		counter++
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
		os.Exit(exitcode)
	})
	// Kills the service after 60 sec
	http.HandleFunc("/timedkill", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This service will be killed in 60 seconds \n")
		fmt.Printf("This service will be killed in 60 seconds \n")
		time.Sleep(time.Duration(sleepTimer) * time.Second)
		os.Exit(exitcode)
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
