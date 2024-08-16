package main

import "fmt"

func main(){

	var a int = 10
	var ptr *int = &a

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)



}