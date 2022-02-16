package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	sleepTimer, _ := strconv.Atoi(os.Getenv("HWC_SLEEPTIMER"))
	sleepCounter := 1

	if sleepTimer == 0 {
		for {
			fmt.Printf("Hello for the %d time \n", sleepCounter)
			time.Sleep(1 * time.Second)
			sleepCounter++
		}
	} else {
		for sleepCounter <= sleepTimer {
			fmt.Printf("Hello for the %d time \n", sleepCounter)
			time.Sleep(1 * time.Second)
			sleepCounter++
		}
	}
}
