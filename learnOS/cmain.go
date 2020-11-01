package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("cat", c)
	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 0; i <= len(thing); i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)

}
