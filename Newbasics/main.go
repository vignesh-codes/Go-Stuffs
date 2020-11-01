package main

import "fmt"

// func main() {
// 	var i = []int{1, 1, 2, 3, 4, 5}

// 	for k, q := range i {
// 		fmt.Printf("%d:%d\n", k, q)
// 	}

// }

func main() {
	var s = map[string]string{"vicky": "gggg", "vic": "fff"}
	for ss, kk := range s {
		fmt.Printf("%s is : %s", ss, kk)
	}
}
