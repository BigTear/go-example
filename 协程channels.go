package main

import "fmt"

func main() {
	// 无缓冲
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	msg := <-message

	fmt.Println(message)
	fmt.Println(msg)

	fmt.Println()
	//通道缓冲
	messageWBuffer := make(chan string, 2)
	messageWBuffer <- "first"
	messageWBuffer <- "second"
	//error
	//messageWBuffer <- "third"

	fmt.Println(<-messageWBuffer)
	fmt.Println(<-messageWBuffer)
}
