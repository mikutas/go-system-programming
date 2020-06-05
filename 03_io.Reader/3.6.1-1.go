package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `1行め
2行め
3行め`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}
}
