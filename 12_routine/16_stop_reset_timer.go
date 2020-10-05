package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("go routine can print.")
	}()
	timer.Stop() // stop timer
	timer.Reset(time.Second) // reset to 1 sec.
}
