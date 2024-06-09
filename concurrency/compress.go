package main

import (
	"fmt"
	"log"
	"slices"
)

func sort(a []int) {
	slices.Sort(a)
}

func main() {

	a1 := []int{4, 3, 2, 1}
	sort(a1)
	log.Println(a1)

	m := make(map[string]int, 12)
	log.Println(m, len(m))

	// bit mac in my case 32 bit
	var i uint = 00
	log.Println(i)

	a := "AABBBBCCCDDDBBBBBCCA" // => A2B4C3D3B5C2A
	// 1234

	counter := 1
	res := ""

	for i := 1; i < len(a); i++ {
		cur := string(a[i])
		prev := string(a[i-1])

		if cur != prev {
			res += fmt.Sprintf("%s%d", prev, counter)
			counter = 1
			continue
		}

		counter++
	}

	res += fmt.Sprintf("%s%d", string(a[len(a)-1]), counter)

	log.Println(res)

}
