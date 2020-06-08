package main

import (
	"io"
	"os"
)

func main() {
	old, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer old.Close()
	new, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer new.Close()

	io.Copy(new, old)
}
