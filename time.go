package main

import (
	"time"
	"fmt"
)

func main() {
	stop := time.After(3 * time.Second)
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			fmt.Printf("hello")
			fmt.Println(5 * time.Minute, 5 + time.Now().Minute())
		case <-stop:
			return
		}
	}
}