package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file := flag.String("file", "listing_0038_many_register_mov", "name of the file inside /listings directory")
	flag.Parse()

	f, err := os.Open("./listings/" + *file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for {
		b := make([]byte, 2)
		n, err := f.Read(b)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if n < 2 {
			break
		}

		fmt.Printf("%#x\n", string(b[:n]))
	}
}
