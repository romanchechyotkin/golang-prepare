package main

import (
	"log"
	"unsafe"
)

type Struct struct {
	//Float64 float64
	Int8    int8   // 8 bite
	Bool    bool   // 1 bite
	String1 string // 0
	String2 string // 0
	Int     int    // 8 bite
}

type Foo struct {
	b int32   // 4 byte
	c [2]bool // 2 byte
	a [2]bool // 2 byte
}

func main() {
	x := &Foo{}
	y := Foo{}

	log.Println(unsafe.Sizeof(x)) // 8
	log.Println(unsafe.Sizeof(y)) // 8

	x1 := &Struct{}
	y1 := Struct{}
	log.Println(unsafe.Sizeof(x1)) // 8
	log.Println(unsafe.Sizeof(y1)) // 64
}
