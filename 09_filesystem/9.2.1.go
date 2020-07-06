package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, _ := os.Create("file.txt")
	a := time.Now()
	f.Write([]byte("緑の怪獣"))
	b := time.Now()
	f.Sync()
	c := time.Now()
	f.Close()
	d := time.Now()
	fmt.Printf("Write: %v\n", b.Sub(a))
	fmt.Printf("Sync: %v\n", c.Sub(b))
	fmt.Printf("Close: %v\n", d.Sub(c))
}
