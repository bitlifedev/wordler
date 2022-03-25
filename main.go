package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
	"wordler/dictionary"
	"wordler/logger"
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
	dict, err := dictionary.Load(DICTIONARY)
	if err != nil {
		fmt.Println(err)
	}

	secret := selectSecretWord(dict)
	fmt.Println("Secret Wordle is: ", secret)

	//Calculate probability in wordPool  and prune unneeded words
	//Select best guess
	//begin loop
	//update stats
	updateStat(&dict)
	//create interface to guess
	//end loop

}

func selectSecretWord(in dictionary.Dictionary) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(in.Words) - 1)
	return in.Words[index].Value
}

//TODO @Justin fix this
func updateStat(dict *dictionary.Dictionary) {
	logger.Log.Printf("%v", dict)
	totalCount := 0.0
	for i := range dict.Map {
		totalCount += float64(dict.Map[i].Count)
	}
	for i := range dict.Words {
		for j := range dict.Words[i].Map {
			dict.Words[i].Map[j] = dictionary.Stats{Count: dict.Words[i].Map[j].Count, P: dict.Words[i].Map[j].P / totalCount}
		}
	}

	fmt.Println(dict)
}
