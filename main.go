package main

import (
	"fmt"
	"runtime"
	"strconv"
	"wordler/dictionary"
)

func init() {
	fmt.Println(runtime.GOOS + " " + runtime.GOARCH + " " +
		strconv.Itoa(runtime.NumCPU()) + " " +
		strconv.Itoa(runtime.MemProfileRate))
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")

	}
}

func main() {
	fmt.Println("Starting Wordler")

	dictionary.LoadDictionary()
}
