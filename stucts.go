package main

import "fmt"

type Person struct{
	age int
	name string
	salary int

}

func main(){

	var _pers Person
	_pers.name = "Mike"
	_pers.age = 9
	_pers.salary = 10

	fmt.Println("name", _pers.name)
	fmt.Println("age", _pers.age)

}