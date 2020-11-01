package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go count("cat")

	go func() {
		count("dog")
		wg.Done()
	}()
	wg.Wait()
}

func count(thing string) {
	for i := 0; i <= len(thing); i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

}
