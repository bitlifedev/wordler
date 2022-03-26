package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
	"wordler/dictionary"
	"wordler/logger"
	"wordler/words"
)

func init() {
}

const (
	VERSION    = "0.1"
	DICTIONARY = "assets/test.dic"
	//DICTIONARY = "assets/English_5_letter_words.dic"
)

func main() {
	fmt.Println("Starting Wordler")
	logger.Log.Printf("Starting Wordler")
	logger.Log.Printf("Server v%s pid=%d started with processes: %d", VERSION, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

	var wordlerDictionary = new(words.WordlerDictionary)
	var err error
	wordlerDictionary, err = dictionary.Load(DICTIONARY)
	if err != nil {
		fmt.Println(err)
	}
	secret := selectSecretWord(wordlerDictionary)
	fmt.Println("Secret Wordle is: ", secret)

	fmt.Println(wordlerDictionary)
}

func selectSecretWord(dict *words.WordlerDictionary) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(dict.Words) - 1)
	return dict.Words[index].Value
}
