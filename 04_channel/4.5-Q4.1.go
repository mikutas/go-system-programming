package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.After(10 * time.Second)
	t := <-timer
	fmt.Println(t)
}
