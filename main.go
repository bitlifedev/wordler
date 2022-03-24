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
	wordPool, err := dictionary.Load("assets/test.dic")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wordPool)

	secret := selectSecretWord(wordPool)
	fmt.Println("Secret Wordle is: ", secret.Word)

	//Calculate probability in wordPool  and prune unneeded words
	//Select best guess
}

func selectSecretWord(in dictionary.Words) dictionary.Word {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(in.Words)
	return in.Words[rand.Intn(max-min+1)+min]
}
