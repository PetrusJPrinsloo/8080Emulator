package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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
	quit := make(chan struct{})
	state := NewState8080(rom, quit)

	// set ticker to run at 2MHz
	// for each tick, run state.Step()
	ticker := time.NewTicker(time.Microsecond * 500)
	for {
		select {
		case <-ticker.C:
			err = state.Step()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
