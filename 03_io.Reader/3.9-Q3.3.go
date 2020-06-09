package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("hoge.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	// zip内にfuga.txtというファイルを作る。中身はwriterに書き込むまで無い
	writer, err := zipWriter.Create("fuga.txt")
	if err != nil {
		panic(err)
	}
	reader := strings.NewReader("Q3.3\n")
	io.Copy(writer, reader)
}
