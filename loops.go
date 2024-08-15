package main

import "fmt"

func main(){

	//for loop

	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	//continue loop
	for i:=0; i<100; i+=10 {

		if(i == 30){
			continue
		}

		fmt.Println(i)
	}

	//break

	for i:=0; i<100; i+=10 {

		if(i == 40){
			break
		}

		fmt.Println(i)
	}

	//range

	fruits := [3] string {"mango","banana","orange"}

	for idx , val := range fruits {
		fmt.Println(idx," ",val)
	}



}