package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./listing_0038_many_register_mov")
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
