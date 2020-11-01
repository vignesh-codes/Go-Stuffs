package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func dieRoll(size int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size)+1
}

func rollTwo(size1, size2 int) (int,int){
	return dieRoll(size1), dieRoll(size2)
}

func Names(input1 string, input2 int) (theResult string, err error){
	theResult = "modified "+ input1 + ", "+ strconv.Itoa(input2)
	return theResult, err
}

func main(){
	fmt.Println("Rolled die of size %d, result is %d/n", 6, dieRoll(6))

	res1, res2 := rollTwo(6,10)

	fmt.Println("Rolled a pair %d, %d .. Results:%d\n",6,10,res1,res2  )
	names,err := Names("Vicky", 44)

	fmt.Println("Names were: %s, value is %v/n", names, err)
}