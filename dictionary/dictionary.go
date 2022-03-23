package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Words struct {
	words []word
	hash  [26]int
}

type word struct {
	word   string
	hash   [26]int
	active bool
}

func open(file string) (*os.File, error) {
	// Open our jsonFile
	contents, err := os.Open(file)
	if err != nil {
		log.Fatalf("Can't open file, %v", file)
	}
	defer func(contents *os.File) {
		if err != nil {
			log.Fatalf("Can't close file, %v", file)
		}
	}(contents)
	fmt.Println("Successfully Opened dictionary")
	return contents, nil
}

func Load(dictionary string) (Words, error) {

	// read our opened xmlFile as a byte array.
	var WS Words
	file, _ := open(dictionary)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		var W word
		W.word = scanner.Text()
		W.hash = hashString(W.word)
		W.active = true
		WS.words = append(WS.words, W)
		for i := range W.hash {
			WS.hash[i] = WS.hash[i] + W.hash[i]
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return WS, nil
}

func hashString(S string) [26]int {

	var myIntArray [26]int
	for i := range S {
		s := strings.ToUpper(string(S[i]))
		index := s[0] - 65
		myIntArray[index]++
	}

	return myIntArray
}
