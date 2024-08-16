package main

import "fmt"

func main(){
	newPrint("Mary",6)
	a,b := printMultiple(4,"wewe ")
	fmt.Println(a,b)

}

func newPrint(name string, age int){
	fmt.Println(name, "name")
	fmt.Println(age,"age")

}

func printMultiple(a int, b string) (result int, txt string){
	result = a + a
	txt = b + "my new string here"
	return
}