// write a generator function which returns a channel
// call a goroutine to return random integer less than or equal to n
package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	number  int
	message string
}

func main() {
	for val := range generate(5) {
		fmt.Println(val)
	}
}

func generate(n int) <-chan node {
	ch := make(chan node)

	go func() {
		defer close(ch)

		for i := 0; i < n; i++ {
			num := rand.Intn(n) + 1

			ch <- node{
				number:  num,
				message: fmt.Sprintf("Your random number is %d", num),
			}
		}
	}()
	return ch
}
