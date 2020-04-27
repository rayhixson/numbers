package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	sequential := flag.Bool("seq", false, "generate sequential 9 digit numbers, default is false - random")
	flag.Parse()

	if *sequential {
		for i := 0; ; i++ {
			fmt.Printf("%09d\n", i)
		}
	} else {
		r := rand.New(rand.NewSource(1))
		for {
			fmt.Printf("%09d\n", r.Int63()/10000000000)
		}
	}
}
