package main

import "fmt"

type person struct{
	name string
	age int
}
func main(){
	p1:=person{name:"muhammed",age: 22 }

fmt.Println("name:",p1.name)
fmt.Println("age:",p1.age)
}