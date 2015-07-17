package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		for t := range ticker.C {
			log.Println(t)
		}
	}()

	s := <-c
	close(c)
	log.Println("Got signal:", s)
}
