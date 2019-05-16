package main

import (
	"fmt"
	"math/rand"
	"time"
)

var
(
	CLOSEA = false
	DATA   = make(map[int]bool)
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	n1 := 100
	n2 := 1000000

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(B, A)
	third(B)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}
