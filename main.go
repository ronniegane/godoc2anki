package main

import (
	"flag"
	"fmt"
)

func main() {
	docURL := flag.String("u", "https://golang.org/pkg", "The URL of the Godoc you wish to parse")
	outFilename := flag.String("o", "test.csv", "The file to output to")

	flag.Parse()

	fmt.Println(*docURL)
	fmt.Println(*outFilename)
}
