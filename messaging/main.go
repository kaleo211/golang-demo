package main

import (
	"fmt"

	"github.com/kaleo211/golang-demo/time"
)

func main() {
	messages := make(chan string)

	go func() {
		time.SleepForSeconds(3)
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
