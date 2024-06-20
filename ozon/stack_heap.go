package main

import "log"

func main() {
	a := answer()
	log.Println(*a)
}

func answer() *int {
	x := 42
	return &x
}
