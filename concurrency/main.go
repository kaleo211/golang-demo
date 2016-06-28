package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/kaleo211/golang-demo/time"
)

func myPrint(content string) {
	for i := 0; i < 3; i++ {
		time.Sleep()

		fmt.Printf("%s: %d\n", content, i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
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
