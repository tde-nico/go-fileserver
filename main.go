package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	var (
		port      string
		directory string
	)

	flag.StringVar(&port, "p", "8777", "The port to connect to")
	flag.StringVar(&directory, "d", ".", "The directory to serve")

	flag.Parse()

	file_server := http.FileServer(http.Dir(directory))
	http.Handle("/", file_server)

	fmt.Printf("[+] Serving %s on 0.0.0.0:%s\n", directory, port)
	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		fmt.Printf("[!] Error: %v\n", err)
	}
}
