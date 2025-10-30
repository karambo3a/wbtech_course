package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// Returns current time using NTP server
// If error occurs prints error message to STDERR and exits with code 1
func GetCurrTime(address string) (time.Time, error) {
	return ntp.Time(address)
}

func main() {
	currTime, err := GetCurrTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ntp error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(currTime)
}
