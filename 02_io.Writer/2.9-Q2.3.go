package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	gzWriter := gzip.NewWriter(w)
	multiWriter := io.MultiWriter(gzWriter, os.Stdout)
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", " ")
	encoder.Encode(source)
	gzWriter.Flush()
	gzWriter.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
