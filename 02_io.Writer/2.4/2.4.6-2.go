package main

import (
	"compress/gzip"
	"encoding/csv"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()

	file2, err := os.Create("test.csv")
	writer2 := csv.NewWriter(file2)
	writer2.Write([]string{"hoge", "fuga"})
	writer2.Flush()
}
