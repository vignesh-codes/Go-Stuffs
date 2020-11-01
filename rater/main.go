package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//Front
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you Name")
	name, _ = reader.ReadString('\n')

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Ratings")
	userRating, _ = reader.ReadString('\n')
	myNum, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Back
	fmt.Printf("%v, %v, %v", name, myNum, time.Now().Format(time.Stamp))

}
