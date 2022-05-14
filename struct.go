package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

func bornPerson(name string) *person {
	return &person{name: name, age: 1, sex: "unkonwn"}
}

func main() {
	fmt.Println(bornPerson("zyd"))
}
