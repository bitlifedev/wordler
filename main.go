package main

import (
	"fmt"
	"os"
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
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Starting Wordler")
	wordPool, err := dictionary.Load("assets/test.dic")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wordPool)
}
