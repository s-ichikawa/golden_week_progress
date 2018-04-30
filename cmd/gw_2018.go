package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	done   = "▓"
	undone = "░"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("Now: 2006-01-02 15:04:05"))

	start_at := time.Date(2018, 4, 28, 0, 0, 0, 0, time.Local)
	end_at := time.Date(2018, 5, 6, 23, 59, 59, 0, time.Local)

	duration := end_at.Sub(start_at)
	passed := now.Sub(start_at)

	passedRate := passed.Hours() / duration.Hours() * 100.0
	if passedRate > 100.0 {
		passedRate = 100.0
	}
	fmt.Printf("Passed: %f%%\n", passedRate)

	bar := strings.Repeat(undone, 20)
	var progress int
	if (passedRate / 5) < 100.0 {
		progress = int(passedRate / 5)
	} else {
		progress = 20
	}

	fmt.Println(strings.Replace(bar, undone, done, progress))

}
