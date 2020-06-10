package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func CopyN(dst io.Writer, src io.Reader, length int64) (written int64, err error) {
	srcLimited := io.LimitReader(src, length)
	written, err = io.Copy(dst, srcLimited)
	return
}

func main() {
	r := strings.NewReader("abcdefghijklmn")
	buf := bytes.NewBufferString("")
	written, err := CopyN(buf, r, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(written)
	fmt.Println(buf)
}
