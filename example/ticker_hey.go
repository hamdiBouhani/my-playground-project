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

func runTicker2() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Println("Hey!")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- struct{}{}
	log.Println("Done")
}
