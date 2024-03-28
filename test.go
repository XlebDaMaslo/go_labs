package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {}

func main() {
	for i := 1; i < 10000; i++ {
		select {

		case m := <-c:
			handle(m)
		case <-time.After(1 * time.Second):
			fmt.Println(i)
		}
	}
}
