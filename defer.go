package main

import (
	"fmt"
	"log"
)

func main() {
	defer fmt.Println("World")
	//panic("Stop")
	//fmt.Println("Hello")

	log.Println(foo())
}

func inc() (i int) {
	defer func() {
		i++
	}()

	log.Println(i)

	return i
}

func foo() (result string) {
	defer func() {
		result = "Change World"
	}()

	return "Hello World"
}
