package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {
	fmt.Println(test1()) // 1
	test3()
	s := "boaT 🛥️."
	fmt.Println("length of string", len(s))
	log.Println(utf8.RuneCountInString(s))

	for i, c := range s {
		fmt.Printf("position %d of '%s'\n", i, string(c))
	}
}

// returns 1
func test1() (result int) {
	defer func() {
		result++
	}()

	return result
}

func test3() {
	var i1 int = 10
	var k = 20
	var i2 *int = &k

	// defer сохраняет значение, которое передается в моменте
	defer printInt("i1", i1)               // 10
	defer printInt("i2 as value", *i2)     // 20
	defer printIntPtr("i2 as pointer", i2) // 2020

	i1 = 1010
	*i2 = 2020
}

func printInt(v string, i int) {
	fmt.Printf("%s=%d\n", v, i)
}

func printIntPtr(v string, i *int) {
	fmt.Printf("%s=%d\n", v, *i)
}
