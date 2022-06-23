package main

import (
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"time"
)

func listenKeyboard(direction chan<- int) {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		switch r {
		case 'a':
			direction <- -1
		case 'd':
			direction <- 1
		}
	}
}

func main() {
	direction := make(chan int, 1)
	go listenKeyboard(direction)
	ticker := time.NewTicker(100 * time.Millisecond)

	location := 5
	for {
		<-ticker.C
		// print spaces to clear the previous location
		for i := 0; i < location; i++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		select {
		case d := <-direction:
			location += d
		default:
		}
		fmt.Println()
	}
}
