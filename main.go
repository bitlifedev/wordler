package main

import (
	"fmt"
	"math/rand"
	"time"
	"wordler/dictionary"
)

func init() {

}

func main() {
	fmt.Println("Starting Wordler")
	dict, err := dictionary.Load("assets/test.dic")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dict)

	secret := selectSecretWord(dict)
	fmt.Println("Secret Wordle is: ", secret)

	//Calculate probability in wordPool  and prune unneeded words
	//Select best guess
}

func selectSecretWord(in dictionary.Dictionary) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(in.Words) - 1)
	return in.Words[index].Value

}
