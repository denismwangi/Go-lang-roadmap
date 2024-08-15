package main

import "fmt"

func main(){

	var testVar = [] int {} //empty slice
	myslice := [] string {"jong","mark","Betty"}

	fmt.Println(testVar)
	fmt.Println(myslice)

	var evenBetter = make([] int,4,9)  //with len
	var evenBetter2 = make([] int,5) //without len

	fmt.Println(evenBetter)
	fmt.Println(evenBetter2)
	fmt.Println(cap(evenBetter))

	evenBetter = append(evenBetter, 5,6,6,6)

	fmt.Println("appended slice", evenBetter)
}