package main

import (
	"context"
	"fmt"
	"os"

	"github.com/billziss-gh/cgofuse/fuse"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
)

type CloudFileSystem struct {
	fuse.FileSystemBase
	bucket *blob.Bucket
}

func main() {
	ctx := context.Background()
	if len(os.Args) < 3 {
		fmt.Println("%s [bucket-path] [mount-point] etc...", os.Args[0])
		os.Exit(1)
	}
	b, err := blob.OpenBucket(ctx, os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer b.Close()
	cf := &CloudFileSystem{bucket: b}
	host := fuse.NewFileSystemHost(cf)
	host.Mount(os.Args[2], os.Args[3:])
}
