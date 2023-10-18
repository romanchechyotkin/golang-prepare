package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	out := compress("AABBBBCCCDDDBBBBBCCA")
	log.Println("out", out)
	if out != "A2B4C3D3B5C2A1" {
		log.Fatal("not equal")
	}

	out = compress("")
	log.Println("out", out)
	if out != "" {
		log.Fatal("not equal")
	}

	out = compress("ABCDE")
	log.Println("out", out)
	if out != "A1B1C1D1E1" {
		log.Fatal("not equal")
	}
}

func compress(in string) string {
	if len(in) == 0 {
		return ""
	}

	var res bytes.Buffer
	//res.WriteString(fmt.Sprintf("%s", string(in[0])))
	counter := 1
	var prevChar byte
	var currChar byte

	for i := 1; i < len(in); i++ {
		prevChar = in[i-1]
		currChar = in[i]

		if prevChar != currChar {
			res.WriteString(fmt.Sprintf("%s%d", string(prevChar), counter))
			counter = 1
		} else {
			counter++
		}
	}

	res.WriteString(fmt.Sprintf("%s%d", string(currChar), counter))

	return res.String()
}
