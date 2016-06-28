package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func myPrint(content string) {
	for i := 0; i < 3; i++ {
		rand.Seed(time.Now().UnixNano())
		duration := time.Duration(rand.Intn(200))
		time.Sleep(duration * time.Millisecond)

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
