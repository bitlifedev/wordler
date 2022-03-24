package main

import (
	"fmt"
	"wordler/dictionary"
)

func init() {

}

func main() {
	fmt.Println("Starting Wordler")
	wordPool, err := dictionary.Load("assets/test.dic")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wordPool)
}
