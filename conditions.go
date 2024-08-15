package main

import "fmt"

func main(){

	a := 5
	b := 6
	c := 8

	if a < 6 {
		fmt.Println("a is not greater")
	}

	if c > 100 {
		fmt.Print("okay men")
	} else if a < b {
		fmt.Println("I knew it!")
	}
}