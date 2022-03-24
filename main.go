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
	VERSION = "0.1"
)

func main() {
	fmt.Println("Starting Wordler")
	logger.Log.Printf("Starting Wordler")
	logger.Log.Printf("Server v%s pid=%d started with processes: %d", VERSION, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))
	dict, err := dictionary.Load("assets/test.dic")
	if err != nil {
		fmt.Println(err)
	}

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

func updateStat() {

}
