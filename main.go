package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		compress             bool
		varName, packageName string
	)

	flag.StringVar(&varName, "var", "bencoded", "name of variable to encode to")
	flag.StringVar(&packageName, "pkg", "main", "name of package")
	flag.BoolVar(&compress, "compress", false, "compress with gzip before encoding")
	flag.Parse()

	r := io.Reader(bufio.NewReader(os.Stdin))
	if compress {
		b := bytes.NewBuffer(make([]byte, 1024))
		w := gzip.NewWriter(b)
		go func(w io.Writer, r io.Reader) {
			if _, err := io.Copy(w, r); err != nil {
				log.Panicf("Error compressing: %v", err)
			}
		}(w, r)
		r = b
	}

	fmt.Printf(`package %s

var %s = []byte{`, packageName, varName)

	first := true
	reading := true
	p := make([]byte, 1024)
	for reading {
		n, err := r.Read(p)
		if err != nil {
			reading = false
			if err != io.EOF {
				log.Printf("Error reading input: %v", err)
			}
		}

		for _, b := range p[:n] {
			if !first {
				fmt.Print(", ")
			}
			first = false
			fmt.Printf("%d", b)
		}
	}
	fmt.Print("}\n")
}
