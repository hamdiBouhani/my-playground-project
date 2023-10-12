package example

import (
	"log"
	"time"
)

func runTicker() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			log.Println("Hey!")
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
}
