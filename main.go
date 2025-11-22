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

	var instructionNames = map[byte]string{
		OpCodeMov: "MOV",
	}

	file := flag.String("file", "listing_0037_single_register_mov", "name of the file inside /listings directory")
	flag.Parse()

	f, err := os.Open("./listings/" + *file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := make([]byte, 2)
	for {
		_, err := f.Read(b)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		if name, ok := instructionNames[b[0]&OpCodeMask]; ok {
			fmt.Println(name)
		} else {
			fmt.Println("Unsupported opcode")
		}
	}
}
