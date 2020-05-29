package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())
	fmt.Fprintf(os.Stdout, "%d\n", 1)
	fmt.Fprintf(os.Stdout, "%s\n", "1.1")
	fmt.Fprintf(os.Stdout, "%f\n", 1.1)
}
