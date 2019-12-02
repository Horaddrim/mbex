package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var target string
var fullMagicByte bool

var magicByteLength int8

func init() {
	flag.BoolVar(&fullMagicByte, "full", false, "Reads the full 16 byte length of the Magic Byte header")
	flag.StringVar(&target, "target", "", "The file to extract the Magic Bytes")
}

// Nasty little function, but I'm in a hurry, don't blame me.
func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	if target == "" {
		fmt.Println("Usage: mbex -target file [-full]")
		os.Exit(1)
	}

	physicalPath, err := filepath.Abs(target)

	check(err)

	fileData, err := os.Open(physicalPath)

	check(err)

	defer fileData.Close()

	if fullMagicByte {
		magicByteLength = 16
	} else {
		magicByteLength = 8
	}

	bytes := make([]byte, magicByteLength)

	fileData.Read(bytes)

	fmt.Printf("Magic Bytes found: %X\n", bytes)
}
