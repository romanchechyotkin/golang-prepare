package main

import "log"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := &Person{
		Name: "Roman",
		Age:  123,
	}
	log.Println(p)
	changeCopy(p)
	log.Println(p)
	changePtr(p)
	log.Println(p)
}

func changeCopy(p *Person) {
	p = &Person{
		Name: "AWED",
		Age:  123213,
	}
}

func changePtr(p *Person) {
	p.Name = "ROMANNNN"
}
