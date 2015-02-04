package main

import (
	"flag"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "port")
	dir := flag.String("dir", ".", "directory to serve from")
	flag.Parse()
	server := http.FileServer(http.Dir(*dir))
	http.Handle("/", server)
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
