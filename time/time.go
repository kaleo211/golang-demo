package time

import (
	"math/rand"
	"time"
)

func Sleep() {
	rand.Seed(time.Now().UnixNano())
	duration := time.Duration(rand.Intn(200))

	time.Sleep(duration * time.Millisecond)
}

func SleepForSeconds(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
}
