package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	const (
		OpCodeMask = 0xFC
		OpCodeMov  = 0x88
	)

	file := flag.String("file", "listing_0037_single_register_mov", "name of the file inside /listings directory")
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

		if b[0]&OpCodeMask == OpCodeMov {
			fmt.Println("Instruction is a MOV")
		}

		fmt.Printf("%#x\n", string(b[:n]))
	}
}
