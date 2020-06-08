package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.Create("random")
	if err != nil {
		panic(err)
	}
	reader := rand.Reader
	limitReader := io.LimitReader(reader, 1024)
	io.Copy(file, limitReader)
}
