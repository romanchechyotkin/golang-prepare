package main

import "log"

func main() {
	if err := handle(); err != nil {
		log.Println(err)
		return
	}
}

func handle() error {
	return &myErr{
		text: "asd",
	}
}

type myErr struct {
	text string
}

func (e *myErr) Error() string {
	return e.text
}
