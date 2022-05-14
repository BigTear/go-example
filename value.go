package main

import "fmt"

func main() {
	whatami := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println("I'm a string")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("I'm an unknown type")
		}
	}
	whatami("Hello")
	whatami(1)
	whatami(true)

}
