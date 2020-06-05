package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行め
2行め
3行め`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
