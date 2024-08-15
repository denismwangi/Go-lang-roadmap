package main

import "fmt"

func main(){ 

	var array_1 = [4] int {1,3,3,4}
	array_2 := [3] int {1,2,3}

	arr1 := [5]int{1:10,2:40}  //initialize partially


	fmt.Println(array_1)
	fmt.Println(array_2)
	fmt.Println(array_1[0])
	fmt.Println(arr1)
	fmt.Println(len(array_1)) //array length
}