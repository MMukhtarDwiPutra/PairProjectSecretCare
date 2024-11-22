package utils

import (
	"fmt"
	"time"
)

func LoadingSpinner(done chan bool) {
	chars := `-\|/`
	for {
		select {
		case <-done:
			return
		default:
			for _, r := range chars {
				fmt.Printf("\rLoading... %c", r)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
}
