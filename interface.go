package main

import "fmt"

type Printer interface{
	Print()
	Display()
}

type Person struct{
	Name string
}

func (p Person) Print(){
	fmt.Println(p.Name)
}

func Display(){
	fmt.Println("Display function implemented")
}

type fileReaderWriter struct {
	filename string
}

func (f fileReaderWriter) readFile(p []byte) (n int, err error){
	//read from file
	return len(p), nil
}

func (f fileReaderWriter) writeFile(p []byte) (n int, err error){
	//write to file
	return len(p), nil
}

func main(){

	p := Person{"John Doe"}
	p.Print()
	Display()

	f := fileReaderWriter{"test.txt"}
	data := make([]byte,1024)
	n, err := f.readFile(data)
	
	if err != nil{
		//handle failure
	}

	_,err = f.writeFile(data)

	if err != nil {
		//handle failure

	}

	fmt.Println(n)

}