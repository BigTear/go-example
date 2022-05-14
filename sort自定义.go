package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (this byLength) Len() int {
	return len(this)
}

func (this byLength) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this byLength) Less(i, j int) bool {
	return len(this[i]) < len(this[j])
}

func main() {
	fruits := byLength{"peach", "banana", "kiwi"}
	sort.Sort(fruits)
	fmt.Println(fruits)
}
