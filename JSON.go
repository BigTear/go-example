package main

import (
	"encoding/json"
	"fmt"
)

type res1 struct {
	Page   int
	Fruits []string
}

type res2 struct {
	Page   int      `json:"Page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
}
