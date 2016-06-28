package main

import (
	"fmt"
	"sync"

	"github.com/kaleo211/golang-demo/time"
)

func myPrint(caller string) {
	for i := 0; i < 3; i++ {
		time.Sleep()

		fmt.Printf("%s: %d\n", caller, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		myPrint("goroutine")
	}()

	myPrint("main")

	// wg.Wait()
	fmt.Println("\nTerminating Program")
}
