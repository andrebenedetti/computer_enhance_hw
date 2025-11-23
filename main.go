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
		DMask      = 0x2
		WMask      = 0x1
		RegMask    = 0x38
		RMMask     = 0x7

		RegAL = 0x0
		RegCL = 0x1
		RegDL = 0x2
		RegBL = 0x3
		RegAH = 0x4
		RegCH = 0x5
		RegDH = 0x6
		RegBH = 0x7
	)

	var instructionNames = map[byte]string{
		OpCodeMov: "mov",
	}

	var regW0Map = map[byte]string{
		RegAL: "al",
		RegCL: "cl",
		RegDL: "dl",
		RegBL: "bl",
		RegAH: "ah",
		RegCH: "ch",
		RegDH: "dh",
		RegBH: "bh",
	}

	var regW1Map = map[byte]string{
		RegAL: "ax",
		RegCL: "cx",
		RegDL: "dx",
		RegBL: "bx",
		RegAH: "sp",
		RegCH: "bp",
		RegDH: "si",
		RegBH: "di",
	}

	file := flag.String("file", "listing_0038_many_register_mov", "name of the file inside /listings directory")
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

		d := (b[0] & DMask) >> 1
		w := b[0] & WMask
		opcode := b[0] & OpCodeMask
		reg := b[1] & RegMask >> 3
		rm := b[1] & RMMask

		if name, ok := instructionNames[opcode]; ok {
			var mapToUse map[byte]string
			if w == 0 {
				mapToUse = regW0Map
			} else {
				mapToUse = regW1Map
			}

			var line string

			if d == 1 {
				line = fmt.Sprintf("%s %s, %s", name, mapToUse[reg], mapToUse[rm])
			} else {
				line = fmt.Sprintf("%s %s, %s", name, mapToUse[rm], mapToUse[reg])
			}

			fmt.Println(line)
		} else {
			fmt.Println("Unsupported opcode")
		}
	}
}
