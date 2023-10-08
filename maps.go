package main

import "log"

type Bar struct {
	Val  int
	Val2 int
}

func main() {
	m := make(map[Bar]int)
	m[Bar{Val: 1, Val2: 1}] = 1
	m[Bar{Val: 2}] = 2
	log.Println(m)
	log.Println(m[Bar{
		Val:  1,
		Val2: 1,
	}])
	log.Println(m[Bar{
		Val:  1,
		Val2: 0,
	}])
	log.Println(m[Bar{
		Val:  2,
		Val2: 0,
	}])

}
