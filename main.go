package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RetrieveROM(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	check(err)
	defer file.Close()

	fileInfo, err := file.Stat()
	check(err)

	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)
	check(err)

	return bytes, err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	rom, err := RetrieveROM(filename)
	check(err)
	state := NewState8080(rom)
	for err == nil {
		fmt.Printf("%d: %02X\n")
		err = state.Step()
	}
}
