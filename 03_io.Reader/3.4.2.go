package main

import (
	"io"
	"os"
)

func main() {
	// go runする場合ファイルのディレクトリに移動
	file, err := os.Open("3.4.2.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
