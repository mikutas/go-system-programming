package main

import (
	"fmt"
	"time"
)

func main() {
	l := NewFileLock("10.2.3.go")
	fmt.Println("try locking...")
	l.Lock()
	fmt.Println("locked!")
	time.Sleep(10 * time.Second)
	l.UnLock()
	fmt.Println("unlock")
}
