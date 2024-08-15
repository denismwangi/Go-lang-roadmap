package main 

import "fmt"

func main(){

	day := 3
	month := 5

	switch day {
		case 1:
			fmt.Println("its monday")
		case 2:
			fmt.Println("its Ty")
		case 3:
			fmt.Println("Gotcha!!")
		default:
			fmt.Println("As always")
	}

	switch month {
		case 1,3,5:
			fmt.Println("odd numbers")
		case 2,4:
			fmt.Println("Even numbers months")
		default:
			fmt.Println("Enda home")
	}
}