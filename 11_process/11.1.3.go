package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	sid, _ := unix.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n", unix.Getpgrp(), sid)
}
