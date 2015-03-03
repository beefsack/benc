package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		varName, packageName string
	)

	flag.StringVar(&varName, "var", "bencoded", "name of variable to encode to")
	flag.StringVar(&packageName, "pkg", "main", "name of package")
	flag.Parse()

	fmt.Printf(`package %s

var %s = []byte{`, packageName, varName)

	first := true
	reading := true
	r := bufio.NewReader(os.Stdin)
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
