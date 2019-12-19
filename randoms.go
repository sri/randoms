package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var totalNums int
var startRange, endRange int
var useSpaces bool

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.IntVar(&totalNums, "n", 1, "total number of random numbers to generate")
	flag.IntVar(&startRange, "start", 1, "start range of the random set (inclusive)")
	flag.IntVar(&endRange, "end", 100, "end range of the random set (inclusive)")
	flag.BoolVar(&useSpaces, "use_spaces", false, "use spaces instead of newlines when outputting")
	flag.Parse()

	for i := 0; i < totalNums; i += 1 {
		val := rand.Intn(endRange-startRange+1) + startRange
		if useSpaces {
			fmt.Print(val)
			if i < totalNums-1 {
				fmt.Print(" ")
			}
		} else {
			fmt.Println(val)
		}
	}
	if useSpaces {
		fmt.Println()
	}
}
