package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseInt("324234", 0, 64)
	fmt.Printf("value: %v \ntype: %T\n", f, f)

	g, _ := strconv.Atoi("12")
	fmt.Printf("value: %v \ntype: %T\n", g, g)

}
